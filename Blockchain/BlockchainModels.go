package Blockchain

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Chain struct {
	Id    string
	Hash  string
	Block string
}

type Transaction struct {
	RandBytes []byte
	PrevBlock string
	Sender    string
	Receiver  string
	Value     uint64
	Signature string
	CurrHash  string
}

type Block struct {
	CurrHash     string                 `json:"currHash"`
	PrevHash     string                 `json:"prevHash"`
	TimeStamp    *timestamppb.Timestamp `json:"timeStamp"`
	Transactions []Transaction          `json:"transactions"`
	BalanceMap   map[string]uint64      `json:"balanceMap"`
	Nonce        uint64                 `json:"nonce"`
	Difficulty   uint64                 `json:"difficulty"`
	Miner        string                 `json:"miner"`
	//Signatures   string                 `json:"signatures"`
	ChainMaster string `json:"chainMaster"`
}

type User struct {
	Uuid        uuid.UUID
	PublicKey   string
	IsUsed      bool
	Affiliation string
}

type Candidate struct {
	Uuid              uuid.UUID
	PublicKey         string
	Description       string
	VotingAffiliation string
}
