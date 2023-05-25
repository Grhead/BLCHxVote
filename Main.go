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
	fmt.Println(Basic.GetBalance("527688a8bb9652449df23fa41e9d667c26c1bebdf8be51ad7bb6f114df7462ee"))
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
	//fmt.Println(Basic.AcceptNewUser("voter1", "hello", "02d3cbd38ba1e5b21eeba1c4985a89c4fd550612640f3f1f9828ebcf19b605a5"))
	fmt.Println(Basic.AcceptLoadUser("02d3cbd38ba1e5b21eeba1c4985a89c4fd550612640f3f1f9828ebcf19b605a5", "4a860dcc29c11ec1c03903b9e6d80c33628946bec853ee4ca5f05a4a8fbca7ae"))
}
