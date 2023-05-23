package main

import (
	"VOX2/Blockchain"
	"fmt"
)

//type BlockHelp struct {
//	Block   *Blockchain.Block `form:"block" json:"block"`
//	Size    uint64            `form:"size" json:"size"`
//	Address string            `form:"address" json:"address"`
//}
//type re struct {
//	AddTxStatus string
//}

func main() {
	db, err := Blockchain.GetFullDb()
	if err != nil {
		return
	}
	for _, v := range db {
		fmt.Println(v)
	}
}
