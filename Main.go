package main

import (
	"fmt"

	bc "BLCHxVote/Blockchain"
)

const (
	DBNAME = "bloch.db"
)

func main() {
	user1 := bc.NewUser()
	user2 := bc.NewUser()
	fmt.Println(user2.Address())
	fmt.Println(user1.Address())
	bc.NewChain(DBNAME, user1.Address())
	chain := bc.LoadChain(DBNAME)

	// for i := 0; i < 3; i++ {
	// 	block := bc.NewBlock(chain.LastHash())
	// 	block.AddTransaction(chain, bc.NewTransaction(user2, "a1", chain.LastHash(), 10))
	// 	block.Accept(chain, user1, make(chan bool))
	// 	chain.AddBlock(block)
	// }
	// block := bc.NewBlock(chain.LastHash())
	// block.AddTransaction(chain, bc.NewTransaction(user1, "a222", chain.LastHash(), 10))
	// block.Accept(chain, user1, make(chan bool))
	// chain.AddBlock(block)
	block := bc.NewBlock(chain.LastHash())
	block.AddTransaction(chain, bc.NewTransaction(user1, user2.Address(), chain.LastHash(), 10))
	block.Accept(chain, user1, make(chan bool))
	chain.AddBlock(block)
	var bl string
	rows, err := chain.DB.Query("SELECT Block FROM BlockChain")
	if err != nil {
		panic("error: query to db")
	}
	for rows.Next() {
		rows.Scan(&bl)
		fmt.Println(bl)
	}
	/*for i := 0; i < 3; i++ {
		block := bc.NewBlock(chain.LastHash())
		block.AddTransaction(chain, bc.NewTransaction(user1, "a1", chain.LastHash(), 10))
		block.Accept(chain, user1, make(chan bool))
		chain.AddBlock(block)
	}
	var bl string
	rows, err := chain.DB.Query("SELECT Block FROM BlockChain")
	if err != nil {
		panic("error: query to db")
	}
	for rows.Next() {
		rows.Scan(&bl)
		fmt.Println(bl)
	}*/
}
