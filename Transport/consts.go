package Transport

import "VOX2/Blockchain"

type ChainHelp struct {
	Master string `form:"master" json:"master"`
	Count  uint64 `form:"count" json:"count"`
}
type MasterHelp struct {
	Master string `form:"master" json:"master"`
}
type UserHelp struct {
	User string `form:"user" json:"user"`
}
type BlockHelp struct {
	Block   *Blockchain.Block `form:"block" json:"block"`
	Size    uint64            `form:"size" json:"size"`
	Address string            `form:"address" json:"address"`
}
type TransactionHelp struct {
	Master string                  `form:"master" json:"master"`
	Tx     *Blockchain.Transaction `form:"transaction" json:"transaction"`
}
type BalanceHelp struct {
	Balance string `form:"balance" json:"balance"`
}
type SizeHelp struct {
	ChainSize string `form:"chainSize" json:"chainSize"`
}