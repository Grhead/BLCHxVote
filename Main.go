package main

import (
	bc "BLCHxVote/Blockchain"
	"fmt"
)

const (
	DBNAME       = "bloch1.db"
	PASSBDNAME   = "Databases/passdb.db"
	PAREBDNAME   = "Databases/paredb.db"
	PUBLICBDNAME = "Databases/pubdb.db"
)

func main() {
	//bc.NewChain(DBNAME)
	chain := bc.LoadChain(DBNAME)
	bc.NewVotePass(PASSBDNAME, PAREBDNAME, PUBLICBDNAME)
	user1 := bc.NewUser(PUBLICBDNAME)
	//user2 := bc.LoadUser("47ad6449aa0885d4598ac42129d1ae789e453aef6ba39cee12c0fd9ee6c0cdc8", PAREBDNAME)
	fmt.Println(chain.Balance("47ad6449aa0885d4598ac42129d1ae789e453aef6ba39cee12c0fd9ee6c0cdc8", chain.Size()))
	fmt.Println(chain.Balance("17f30ab09a5ff197c1c0b381f7b6412ff4048520f3a949e2774ae2fe97019266", chain.Size()))

	//fmt.Println(user1.Address())
	//fmt.Println(user2.Address())
	//bc.GetTokens(user2, chain, 50)
	bc.Purse("ASD", PASSBDNAME)
	bc.Private("ASD", "Hello", PASSBDNAME, PAREBDNAME, user1.Address())

	//block := bc.NewBlock(chain.LastHash())
	//block.AddTransaction(chain, bc.NewTransaction(user1, user2.Address(), chain.LastHash(), 10))
	//block.Accept(chain)
	//chain.AddBlock(block)
	/*var bl string
	rows, err := chain.DB.Query("SELECT Block FROM BlockChain")
	if err != nil {
		panic("error: query to db")
	}
	for rows.Next() {
		rows.Scan(&bl)
		fmt.Println(bl)
	}*/
}
