package main

import (
	"VOX2/Node"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("LowConf/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	Node.Qwe()
	//_, err = Blockchain.NewChain(100, "first")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//balance, err := Blockchain.Balance("first", "first")
	//fmt.Println(balance)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//_, err = Blockchain.NewChain(50, "second")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for i := 0; i < 100; i++ {
	//item, err := Blockchain.NewPublicKeyItem("first")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//lh, err := Blockchain.LastHash("first")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//newTx, err := Blockchain.NewTransactionFromChain("first", item.Address(), lh, 1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//block, err := Blockchain.NewBlock(lh, "test", "first")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = block.AddTransaction(newTx, "first")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//t, err := item.Private()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//y, err := Blockchain.LoadToEnterAlreadyUser(t)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = block.Accept(y, "first", make(chan bool))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = Blockchain.AddBlock(block)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//}

	//TODO release server
	//l, err := net.Listen("tcp", viper.GetString("PORT"))
	//if err != nil {
	//	log.Fatal(err)
	//}
}
