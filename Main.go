package main

import (
	"VOX2/Basic"
	"fmt"
)

func main() {
	//candidates, err := Basic.CallViewCandidates()
	//var Candidates, _ = Basic.CallViewCandidates()
	//for _, v := range Candidates {
	//	//fmt.Printf("%x\n", v.Uuid.String())
	//	fmt.Println(v)
	//}
	//fmt.Println(Basic.GetBalance("527688a8bb9652449df23fa41e9d667c26c1bebdf8be51ad7bb6f114df7462ee"))
	//fmt.Println(Basic.ChainSize("water"))
	//chain, _ := Basic.GetPartOfChain("water")
	//for _, v := range chain {
	//	fmt.Printf("%v\n", v)
	//}
	//chain, _ := Basic.GetFullChain()
	//for _, v := range chain {
	//	fmt.Printf("%v\n", v)
	//}
	//_, err := Basic.CallCreateVoters("voter1", "water")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, v := range voters {
	//	fmt.Printf("%v\n", v)
	//}
	//fmt.Println(Basic.AcceptNewUser("voter1", "hello", "18aa8c378b646d706aac848373851c8eb666e676dff461aeef97b0861c4d45d3"))
	//fmt.Println(Basic.AcceptLoadUser("18aa8c378b646d706aac848373851c8eb666e676dff461aeef97b0861c4d45d3", "6378bd543591aa66bab5804baa60d372221a49c24658127a20f634ea75726dd5"))
	//
	block, err := Basic.ChainTXBlock("18aa8c378b646d706aac848373851c8eb666e676dff461aeef97b0861c4d45d3", "water", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(block)

	//addresses, err := Basic.ReadAddresses()
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//client := req.C().DevMode()

	//s1 := gocron.NewScheduler(time.UTC)
	//s1.Every(1).Seconds().Do(task)
	//s1.StartAsync()
	//time.Sleep(time.Second * 10)
}

func task() {
	fmt.Println("I am running task.")
}
