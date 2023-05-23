package main

import (
	"VOX2/Basic"
	"fmt"
)

func main() {
	//candidates, err := Basic.CallViewCandidates()
	var Candidates, _ = Basic.CallViewCandidates()
	for _, v := range Candidates {
		//fmt.Printf("%x\n", v.Uuid.String())
		fmt.Println(v)
	}
}
