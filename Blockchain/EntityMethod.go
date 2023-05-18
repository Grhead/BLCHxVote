package Blockchain

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sort"
)

func (user *User) Address() string {
	return user.PublicKey
}

func (user *User) Private() (string, error) {
	var PrivateKey string
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return "", err
	}
	db.Raw("SELECT PrivateKey FROM KeyLinks WHERE PublicKey = ?", user.Address()).Scan(&PrivateKey)
	if PrivateKey == "" {
		return "", errors.New("public key does not exist")
	}
	return PrivateKey, nil
}

func (tran *Transaction) Hash() string {
	return HashSum(string(tran.RandBytes) + tran.Sender + tran.PrevBlock + tran.Receiver + string(ToBytes(tran.Value)))
}

func (block *Block) Hash() string {
	var tempHash string
	for _, tx := range block.Transactions {
		tempHash = HashSum(tempHash +
			tx.CurrHash)
	}
	var list []string
	for hash := range block.BalanceMap {
		list = append(list, hash)
	}
	sort.Strings(list)
	for _, hash := range list {
		tempHash = HashSum(tempHash +
			hash +
			string(ToBytes(block.BalanceMap[hash])))
	}

	return HashSum(
		tempHash +
			string(ToBytes(block.Difficulty)) +
			block.PrevHash +
			block.Miner +
			block.TimeStamp.AsTime().String())
}

func (tran *Transaction) Sign(privateKey string) string {
	return Sign(privateKey, tran.CurrHash)
}

func (block *Block) Sign(privateKey string) string {
	return Sign(privateKey, block.CurrHash)
}

func (block *Block) Proof(ch chan bool) uint64 {
	qwe := ProofOfWork(block.CurrHash, uint8(block.Difficulty), ch)
	fmt.Println("==========> ", qwe)
	return qwe
}

func (block *Block) Accept(user *User, master string, ch chan bool) error {
	curTime, err := GetTime()
	if err != nil {
		return err
	}
	block.TimeStamp = curTime
	block.CurrHash = block.Hash()
	//privateKey, err := user.Private()
	if err != nil {
		return err
	}
	//block.Signatures = block.Sign(privateKey)
	block.Nonce = block.Proof(ch)
	//block.ChainMaster = master
	return nil
}

func (block *Block) AddTransaction(tran *Transaction) error {
	if tran == nil {
		return errors.New("tx = null")
	}
	if tran.Value == 0 {
		return errors.New("tx value = 0")
	}
	if len(block.Transactions) == TxsLimit {
		return errors.New("len tx = limit")
	}
	var err error
	var balanceInChain uint64
	balanceInTX := tran.Value
	if value, ok := block.BalanceMap[tran.Sender]; ok {
		balanceInChain = value
	} else {
		balanceInChain, err = Balance(tran.Sender, block.ChainMaster)
		if err != nil {
			return err
		}
	}
	if balanceInTX > balanceInChain {
		return errors.New("not enough funds")
	}

	block.BalanceMap[tran.Sender] = balanceInChain - balanceInTX
	err = block.AddBalance(tran.Receiver, tran.Value, block.ChainMaster)
	if err != nil {
		return err
	}
	block.Transactions = append(block.Transactions, *tran)
	return nil
}

func (block *Block) AddBalance(receiver string, value uint64, master string) error {
	var balanceInChain uint64
	var err error
	if v, ok := block.BalanceMap[receiver]; ok {
		balanceInChain = v
	} else {
		balanceInChain, err = Balance(receiver, master)
		if err != nil {
			return err
		}
	}
	block.BalanceMap[receiver] = balanceInChain + value
	return nil
}
