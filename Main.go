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
	//_, _ = Basic.PrintBalance("527688a8bb9652449df23fa41e9d667c26c1bebdf8be51ad7bb6f114df7462ee")
	//fmt.Println(Basic.ChainSize("water"))
	//chain, _ := Basic.GetPartOfChain("water")
	//for _, v := range chain {
	//	fmt.Printf("%v\n", v)
	//}
	//chain, _ := Basic.GetFullChain()
	//for _, v := range chain {
	//	fmt.Printf("%v\n", v)
	//}
	_, err := Basic.CallCreateVoters("voter1", "water")
	if err != nil {
		fmt.Println(err)
	}
	//for _, v := range voters {
	//	fmt.Printf("%v\n", v)
	//}
	fmt.Println(Basic.AcceptNewUser("voter1", "hello", "a4ec71ba4f871f739ac2ca565a16a24577c3d29823b8b0d245124b6b5ec79ea0"))
}
