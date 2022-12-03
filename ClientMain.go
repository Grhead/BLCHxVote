package main

import (
	bc "BLCHxVote/Blockchain"
	nt "BLCHxVote/Network"
	"encoding/json"
	"fmt"
)

func main() {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	printBalance("c470ba4dbd73f5378988f2e89d02449c39aa6be6b26cc3ae7f01b6fe6c9c76bf")
	chainSize()
	chainTXBlock("c470ba4dbd73f5378988f2e89d02449c39aa6be6b26cc3ae7f01b6fe6c9c76bf", 1, "ASD", "Databases/passdb.db")
	//User = bc.LoadUser("47ad6449aa0885d4598ac42129d1ae789e453aef6ba39cee12c0fd9ee6c0cdc8", "Databases/paredb.db")
	//chainTX("17f30ab09a5ff197c1c0b381f7b6412ff4048520f3a949e2774ae2fe97019266", 1, "ASD", "Databases/passdb.db")
	printBalance("c470ba4dbd73f5378988f2e89d02449c39aa6be6b26cc3ae7f01b6fe6c9c76bf")
	chainSize()
}
func printBalance(useraddr string) {
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_BLNCE,
			Data:   useraddr,
		})
		if res == nil {
			continue
		}
		fmt.Printf("Balance (%s): %s coins\n", addr, res.Data)
	}
	fmt.Println()
}
func chainSize() {
	res := nt.Send(Addresses[0], &nt.Package{
		Option: GET_CSIZE,
	})
	if res == nil || res.Data == "" {
		fmt.Println("failed: getSize\n")
		return
	}
	fmt.Printf("Size: %s blocks\n\n", res.Data)
}

/*func chainBlock(splited string) {
	num, err := strconv.Atoi(splited)
	if err != nil {
		fmt.Println("failed: strconv.Atoi(num)\n")
		return
	}
	res := nt.Send(Addresses[0], &nt.Package{
		Option: GET_BLOCK,
		Data:   fmt.Sprintf("%d", num-1),
	})
	if res == nil || res.Data == "" {
		fmt.Println("failed: getBlock\n")
		return
	}
	fmt.Printf("[%d] => %s\n", num, res.Data)
}
*/
/*
	func chainBalance(splited []string) {
		if len(splited) != 2 {
			fmt.Println("fail: len(splited) != 2\n")
			return
		}
		printBalance(splited[1])
	}
*/
func chainPrint() {
	for i := 0; ; i++ {
		res := nt.Send(Addresses[0], &nt.Package{
			Option: GET_BLOCK,
			Data:   fmt.Sprintf("%d", i),
		})
		if res == nil || res.Data == "" {
			break
		}
		fmt.Printf("[%d] => %s\n", i+1, res.Data)
	}
	fmt.Println()
}

func chainTX(candidate string, num uint64, datapass string, filename string) {
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_LHASH,
		})
		if res == nil {
			continue
		}
		tx := bc.NewTransaction(User, candidate, bc.Base64Decode(res.Data), num, datapass, filename)
		res = nt.Send(addr, &nt.Package{
			Option: ADD_TRNSX,
			Data:   bc.SerializeTX(tx),
		})
		if res == nil {
			continue
		}
		if res.Data == "ok" {
			fmt.Printf("ok: (%s)\n", addr)
		} else {
			fmt.Printf("fail: (%s)\n", addr)
		}
	}
	fmt.Println()
}
func chainTXBlock(john string, num uint64, datapass string, filename string) {
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_LHASH,
		})
		if res == nil {
			continue
		}
		tx := bc.NewTransactionBlock(john, bc.Base64Decode(res.Data), num, datapass, filename)
		res = nt.Send(addr, &nt.Package{
			Option: ADD_TRNSX,
			Data:   bc.SerializeTX(tx),
		})
		if res == nil {
			continue
		}
		if res.Data == "ok" {
			fmt.Printf("ok: (%s)\n", addr)
		} else {
			fmt.Printf("fail: (%s)\n", addr)
		}
	}
	fmt.Println()
}
