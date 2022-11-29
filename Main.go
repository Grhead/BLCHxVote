package main

import (
	bc "BLCHxVote/Blockchain"
	"fmt"
)

const (
	DBNAME       = "bloch.db"
	PASSBDNAME   = "passdb.db"
	PAREBDNAME   = "paredb.db"
	PUBLICBDNAME = "pubdb.db"
)

func main() {
	bc.NewChain(DBNAME)
	chain := bc.LoadChain(DBNAME)
	bc.NewVotePass(PASSBDNAME, PAREBDNAME, PUBLICBDNAME)
	user1 := bc.NewUser(PUBLICBDNAME)
	user2 := bc.LoadUser("47ad6449aa0885d4598ac42129d1ae789e453aef6ba39cee12c0fd9ee6c0cdc8", PAREBDNAME)
	fmt.Println(user1.Address())
	fmt.Println(user2.Address())
	bc.GetTokens(user1, chain, 20)
	bc.Purse("ASD", PASSBDNAME)
	bc.Private("ASD", "Hello", PASSBDNAME, PAREBDNAME, user1.Address())

	//block := bc.NewBlock(chain.LastHash())
	//block.AddTransaction(chain, bc.NewTransaction(user1, user2.Address(), chain.LastHash(), 10))
	//block.Accept(chain)
	//chain.AddBlock(block)
	//var bl string
	//rows, err := chain.DB.Query("SELECT Block FROM BlockChain")
	//if err != nil {
	//	panic("error: query to db")
	//}
	//for rows.Next() {
	//	rows.Scan(&bl)
	//	fmt.Println(bl)
	//}
}
