package database

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"crypto/sha256"
	"io/ioutil"
)

type SnapShot [32]byte

type State struct {
	Balances map[Account]uint
	txMempool []Tx
	dbFile *os.File
	snapshot SnapShot
}

func (s *State) doSnapShot() error {
	_, err := s.dbFile.Seek(0,0)
	if err!=nil {
		return err
	}

	txsData, err := ioutil.ReadAll(s.dbFile)
	if err != nil {
		return nil
	}
	s.snapshot = sha256.Sum256(txsData)
	return nil
}

func NewStateFromDisk() (*State, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	gen, err := loadGenesis(filepath.Join(cwd, "database","genesis.json"))
	if err != nil {
		return nil, err
	}

	balances := make(map[Account]uint)
	for account, balance := range gen.Balances {
		balances[account] = balance
	}

	txDbFilePath := filepath.Join(cwd,"database","tx.db")
	f, err := os.OpenFile(txDbFilePath,os.O_APPEND|os.O_RDWR,0600)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	state := &State{balances,make([]Tx,0),f,SnapShot{}}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var tx Tx
		err = json.Unmarshal(scanner.Bytes(), &tx)
		if err != nil {
			return nil, err
		}

		if err := state.apply(tx); err != nil {
			return nil, err
		}
	}

	err = state.doSnapShot()
	if err != nil {
		return nil, err
	}

	return state, nil
}

func (s *State) Add(tx Tx)  error {
	if err := s.apply(tx); err != nil {
		return err
	}

	s.txMempool = append(s.txMempool, tx)
	return nil
}

func (s *State) apply( tx Tx) error {
	if tx.isReward() {
		s.Balances[tx.To] += tx.Value
		return nil
	}

	if tx.Value > s.Balances[tx.From] {
		return fmt.Errorf("insufficient balance")
	}

	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value

	return nil
}

func (s *State) Persist() (SnapShot , error) {
	mempool := make([]Tx, len(s.txMempool))
	copy(mempool,s.txMempool)

	for i:=0; i < len(mempool); i++ {
		txJson, err := json.Marshal(mempool[i])
		if err!=nil {
			return SnapShot{},nil
		}

		fmt.Printf("Presisting new Transaction to disk\n")
		fmt.Printf("\t%s\n",txJson)

		if _, err = s.dbFile.Write(append(txJson,'\n')); err != nil {
			return SnapShot{},nil
		}

		err = s.doSnapShot()
		if err!=nil {
			return SnapShot{},err
		}
		fmt.Printf("New DB snapshot is : %x\n",s.snapshot)

		s.txMempool = s.txMempool[i:]
	}
	return s.snapshot, nil
}

func (s *State) Close() error{
	return s.dbFile.Close()
}

func (s *State) LatestSnapshot() SnapShot {
	return s.snapshot
}