package main

import (
	"VOX2/Blockchain"
	"fmt"
)

type BlockHelp struct {
	Block   *Blockchain.Block `form:"block" json:"block"`
	Size    uint64            `form:"size" json:"size"`
	Address string            `form:"address" json:"address"`
}
type re struct {
	AddTxStatus string
}

func main() {
	db, err := Blockchain.GetFullDb()
	if err != nil {
		return
	}
	for _, v := range db {
		fmt.Println(v)
	}
	//master := "bulk"
	//var blocksResponse []*Blockchain.Block
	//client := req.C().DevMode()
	//_, err := client.R().
	//	SetBody(&master).
	//	SetSuccessResult(&blocksResponse).
	//	Post(fmt.Sprintf("http://%s/getblock", strings.Trim("localhost:9595", "\"")))
	//if err != nil {
	//	panic(err)
	//}
	/*viper.SetConfigFile("LowConf/config.env")
	err9 := viper.ReadInConfig()
	if err9 != nil {
		fmt.Println(err9)
		return
	}*/
	//item11, err := Blockchain.NewPublicKeyItem("motor")
	//err = Blockchain.NewDormantUser("pass1333")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//_, err = Blockchain.RegisterGeneratePrivate("pass1333", "salt1", item11.Address())
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//tx1, err := Blockchain.NewTransactionFromChain("motor", item11, 100)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//
	//lh, err := Blockchain.LastHash("water")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//block, err := Blockchain.NewBlock(lh, "water")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//err = block.AddTransaction(tx1)
	//if err != nil {
	//	return
	//}
	//err = block.Accept(make(chan bool))
	//if err != nil {
	//	return
	//}
	//var qwe BlockHelp
	//qwe.Block = block
	//i, _ := Blockchain.Size("water")
	//qwe.Size = i
	//qwe.Address = "9595"
	////client := req.C().DevMode()
	//var result re
	//client := req.C().DevMode()
	//resp, err := client.R().
	//	SetBody(&qwe).
	//	SetSuccessResult(&result).
	//	Post("http://localhost:8585/addblock")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if !resp.IsSuccessState() {
	//	fmt.Println("bad response status:", resp.Status)
	//	return
	//}
	//balance, err := Basic.PrintBalance("895838c26839afb0b9998236edf80ad335f4d5607c3e709b019193d9e6c55cde")
	//if err != nil {
	//	fmt.Println(balance, err)
	//	panic(err)
	//}
	//fmt.Println(balance, err)
	//dengi, err := Blockchain.Balance("895838c26839afb0b9998236edf80ad335f4d5607c3e709b019193d9e6c55cde")
	//if err != nil {
	//	return
	//}
	//fmt.Println(dengi)
	//fmt.Println(balance)
	/*item11, err := Blockchain.NewPublicKeyItem("motor")
	err = Blockchain.NewDormantUser("pass1")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = Blockchain.RegisterGeneratePrivate("pass1", "salt1", item11.Address())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	item12, err := Blockchain.NewPublicKeyItem("motor")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx1, err := Blockchain.NewTransactionFromChain("motor", item11, 1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = Blockchain.NewDormantUser("pass2")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = Blockchain.RegisterGeneratePrivate("pass2", "salt1", item12.Address())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	item13, err := Blockchain.NewPublicKeyItem("motor")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = Blockchain.NewDormantUser("pass3")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = Blockchain.RegisterGeneratePrivate("pass3", "salt1", item13.Address())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	item21, err := Blockchain.NewPublicKeyItem("water")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = Blockchain.NewDormantUser("pass4")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = Blockchain.RegisterGeneratePrivate("pass4", "salt1", item21.Address())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	item22, err := Blockchain.NewPublicKeyItem("water")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = Blockchain.NewDormantUser("pass5")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = Blockchain.RegisterGeneratePrivate("pass5", "salt1", item22.Address())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	item31, err := Blockchain.NewPublicKeyItem("bulk")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = Blockchain.NewDormantUser("pass6")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = Blockchain.RegisterGeneratePrivate("pass6", "salt1", item31.Address())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	_, err = Blockchain.NewChain(500, "motor")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = Blockchain.NewChain(100, "water")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = Blockchain.NewChain(10, "bulk")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx1, err := Blockchain.NewTransactionFromChain("motor", item11, 1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx2, err := Blockchain.NewTransactionFromChain("motor", item12, 1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx3, err := Blockchain.NewTransactionFromChain("motor", item13, 1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx4, err := Blockchain.NewTransactionFromChain("water", item21, 1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx5, err := Blockchain.NewTransactionFromChain("water", item22, 1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx6, err := Blockchain.NewTransactionFromChain("bulk", item31, 1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx7, err := Blockchain.NewTransactionFromChain("motor", item12, 1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	lh, err := Blockchain.LastHash("motor")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	block, err := Blockchain.NewBlock(lh, "motor")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block.AddTransaction(tx1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block.AddTransaction(tx2)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block.AddTransaction(tx3)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block.AddTransaction(tx7)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	lh, err = Blockchain.LastHash("water")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	block1, err := Blockchain.NewBlock(lh, "water")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block1.AddTransaction(tx4)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block1.AddTransaction(tx5)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	lh, err = Blockchain.LastHash("bulk")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	block2, err := Blockchain.NewBlock(lh, "bulk")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block2.AddTransaction(tx6)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block.Accept(make(chan bool))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = Blockchain.AddBlock(block)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block1.Accept(make(chan bool))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = Blockchain.AddBlock(block1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = block2.Accept(make(chan bool))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = Blockchain.AddBlock(block2)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}*/
	//TODO release server
	//l, err := net.Listen("tcp", viper.GetString("PORT"))
	//if err != nil {
	//	log.Fatal(err)
	//}
}
