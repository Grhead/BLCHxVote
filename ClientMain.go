package main

import (
	bc "BLCHxVote/Blockchain"
	nt "BLCHxVote/Network"
	"database/sql"
	"encoding/json"
	"fmt"
)

func main() {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	//printBalance("7921b2bb7c20ad655e713b3bbedd3a91ad65c114a63e6dd32d74632d59d7b98c")
	//chainSize()
	//User = bc.LoadUser("47ad6449aa0885d4598ac42129d1ae789e453aef6ba39cee12c0fd9ee6c0cdc8", PAREDBNAME)
	//chainTXBlock("7921b2bb7c20ad655e713b3bbedd3a91ad65c114a63e6dd32d74632d59d7b98c", 1, "ASD", PASSDBNAME)
	////chainTX("7921b2bb7c20ad655e713b3bbedd3a91ad65c114a63e6dd32d74632d59d7b98c", 1, "ASD", PASSDBNAME)
	//printBalance("7921b2bb7c20ad655e713b3bbedd3a91ad65c114a63e6dd32d74632d59d7b98c")
	//chainSize()
	chainPrint()
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

func ViewCandidates(filename string) []string {
	db, _ := sql.Open("sqlite3", filename)
	rows, _ := db.Query("SELECT * FROM CandidateDB")
	defer db.Close()
	var temp string
	var results []string
	for rows.Next() {
		rows.Scan(&temp)
		results = append(results, temp)
	}
	return results
}

func chainTX(candidate string, num uint64, datapass string, filename string) bool {
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
			return true
		} else {
			fmt.Printf("fail: (%s)\n", addr)
			return false
		}
	}
	fmt.Println()
	return false
}
func chainTXBlock(john string, num uint64, datapass string, filename string) bool {
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
			return true
		} else {
			fmt.Printf("fail: (%s)\n", addr)
			return false
		}
	}
	fmt.Println()
	return false
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
}*/
/*
	func chainBalance(splited []string) {
		if len(splited) != 2 {
			fmt.Println("fail: len(splited) != 2\n")
			return
		}
		printBalance(splited[1])
	}
*/
