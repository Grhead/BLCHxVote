package main

import (
	"VOX2/Blockchain"
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

	err = Blockchain.NewChain(100, "hfheffe")
	h, _ := Blockchain.LastHash("hfheffe")
	fmt.Println(h)
	block, err := Blockchain.NewBlock(h, "user1", "hfheffe")
	block.CurrHash = block.Hash()
	err = Blockchain.AddBlock(block)
	//TODO release server
	//Blockchain.NewChain(100, "first")
	//Blockchain.NewPublicKeyItem("first")
	//l, err := net.Listen("tcp", viper.GetString("PORT"))
	//if err != nil {
	//	log.Fatal(err)
	//}
}
