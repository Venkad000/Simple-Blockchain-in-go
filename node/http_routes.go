package node

import (
	"blockchain/database"
	"net/http"
)

type ErrRes struct {
	Error string
}

type BalancesRes struct {
	Hash 	database.Hash
	Balances map[database.Account]uint
}

type TxAddReq struct {
	From string
	To string
	Value uint
	Data string
}

type TxAddRes struct {
	Hash database.Hash
}

type StatusRes struct {
	Hash database.Hash
	Number uint64
	KnownPeers []PeerNode
}


func listBalancesHandler(w http.ResponseWriter,r *http.Request, state *database.State) {
	writeRes(w,BalancesRes{state.LatestBlockHash(),state.Balances})
}

func statusHandler(w http.ResponseWriter, r* http.Request, n *Node) {
	res := StatusRes {
		Hash : n.state.LatestBlockHash(),
		Number: n.state.LatestBlock().Header.Number,
		KnownPeers: n.knownPeers,
	}

	writeRes(w, res)
}

func txAddHandler(w http.ResponseWriter, r* http.Request, state *database.State) {
	req := TxAddReq{}
	err := readReq(r, &req)
	if err != nil {
		writeErrRes(w, err)
		return
	}

	tx := database.NewTx(database.NewAccount(req.From), database.NewAccount(req.To), req.Value, req.Data)

	err = state.AddTx(tx)
	if err != nil {
		writeErrRes(w, err)
		return
	}

	hash, err := state.Persist()
	if err != nil {
		writeErrRes(w, err)
		return
	}
	
	writeRes(w, TxAddRes{hash})
}
