package Contract

import (
	pr "BLCHxVote/API/Proto"
	bc "BLCHxVote/Blockchain"
	nt "BLCHxVote/Network"
	"context"
	"encoding/json"
	"fmt"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	"io/ioutil"
)

var (
	Addresses []string
)
var User *bc.User

func init() {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
}
func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

const (
	ADD_BLOCK = iota + 1
	ADD_TRNSX
	GET_BLOCK
	GET_LHASH
	GET_BLNCE
	GET_CSIZE
)

type GRserver struct {
	pr.BLCH_ContractServer
}

func (s *GRserver) mustEmbedUnimplementedBLCH_ContractServer() {
	panic("implement me")
}

func (s *GRserver) ChainSize(context.Context, *pr.Wpar) (*pr.ResponseSize, error) {
	res := nt.Send(Addresses[0], &nt.Package{
		Option: GET_CSIZE,
	})
	fmt.Println(res.Data)
	if res == nil || res.Data == "" {
		//fmt.Println("failed: getSize\n")
		return &pr.ResponseSize{Size: "fail"}, nil
	}
	//fmt.Printf("Size: %s blocks\n\n", res.Data)
	srr := fmt.Sprintf("%s", res.Data)
	return &pr.ResponseSize{Size: srr}, nil

}

/*
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
*/
/*func printBalance(useraddr string) {
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
}*/
/*func ChainSize() string {
	res := nt.Send(Addresses[0], &nt.Package{
		Option: GET_CSIZE,
	})
	fmt.Println(res.Data)
	if res == nil || res.Data == "" {
		//fmt.Println("failed: getSize\n")
		return "fail"
	}
	//fmt.Printf("Size: %s blocks\n\n", res.Data)
	srr := fmt.Sprintf("%s", res.Data)
	return srr
}*/

/*func chainPrint() {
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
}*/

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
