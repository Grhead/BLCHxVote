package main

import (
	nt "BLCHxVote/Network"
	//bc "BLCHxVote/Blockchain"
	"fmt"
)

var (
	Addresses1 []string
)

const (
	ADD_BLOCK1 = iota + 1
	GET_BLOCK1
)

func main() {
	Addresses1 = append(Addresses1, ":8080")
	chainPrint()
}
func chainPrint() {
	for i := 0; ; i++ {
		res := nt.Send(Addresses1[0], &nt.Package{
			Option: GET_BLOCK1,
			Data:   fmt.Sprintf("%d", i),
		})
		if res == nil || res.Data == "" {
			break
		}
		fmt.Printf("[%d] => %s\n", i+1, res.Data)
	}
	fmt.Println()
}
