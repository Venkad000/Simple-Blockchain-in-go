package database

type Account string

type Tx struct {
	From Account
	To Account
	Value uint
	Data string
}

func NewAccount(value string) Account {
	return Account(value)
}

func (t Tx) isReward() bool {
	return t.Data == "reward"
}

func NewTx(from Account, to Account, value uint, data string) Tx {
	return Tx{from, to,value,data}
}