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
	User      *bc.User
)

const (
	ADD_BLOCK = iota + 1
	ADD_TRNSX
	GET_BLOCK
	GET_LHASH
	GET_BLNCE
	GET_CSIZE
	PASSDBNAME      = "Databases/passdb.db"
	PAREDBNAME      = "Databases/paredb.db"
	PUBLICDBNAME    = "Databases/pubdb.db"
	CANDIDATEDBNAME = "Databases/candidate.db"
)

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
} /*
func (s *GRserver) TimeBlock(ctx context.Context, ld *pr.BlockData) (*pr.BlockData, error) {
	var srr string
	num, err := strconv.Atoi(ld.BlockNum)
	if err != nil {
		srr = "TimeInvalidBlock"
	}
	res := nt.Send(Addresses[0], &nt.Package{
		Option: GET_BLOCK,
		Data:   fmt.Sprintf("%d", num-1),
	})
	if res == nil || res.Data == "" {
		srr = "TimeInvalidResData"
	}
	//fmt.Printf("[%d] => %s\n", num, res.Data)
	srr = fmt.Sprintf("[%d] => %s\n", num, res.Data)
	return &pr.BlockData{BlockNum: srr}, nil
}
func (s *GRserver) Balance(ctx context.Context, address *pr.Address) (*pr.Lanb, error) {
	var srr string
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_BLNCE,
			Data:   address.Useradrr,
		})
		if res == nil {
			continue
		}
		srr = fmt.Sprintf("%s, %s", addr, res.Data)
	}
	return &pr.Lanb{Balance: srr}, nil
}
func (s *GRserver) ViewCandidates(context.Context, *pr.Wpar) (*pr.CandidateList, error) {
	db, _ := sql.Open("sqlite3", CANDIDATEDBNAME)
	rows, _ := db.Query("SELECT * FROM CandidateDB")
	defer db.Close()
	var temp string
	fmt.Println(rows)
	var results []string
	for rows.Next() {
		rows.Scan(&temp)
		fmt.Println(temp)
		results = append(results, temp)
	}
	return &pr.CandidateList{Candidate: results}, nil
}
func (s *GRserver) Transfer(ctx context.Context, ld *pr.LowDataChain) (*pr.IsComplited, error) {
	var srr bool
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_LHASH,
		})
		fmt.Println("YES")
		if res == nil {
			continue
		}
		fmt.Println("YES1")
		tx := bc.NewTransactionBlock(ld.UserCandidate, bc.Base64Decode(res.Data), ld.Num)
		fmt.Println("YES2")
		res = nt.Send(addr, &nt.Package{
			Option: ADD_TRNSX,
			Data:   bc.SerializeTX(tx),
		})
		fmt.Println("YES3")
		if res == nil {
			continue
		}
		fmt.Println("YES4")
		if res.Data == "ok" {
			srr = true
		} else {
			srr = false
		}
	}
	return &pr.IsComplited{Ic: srr}, nil
}
func (s *GRserver) Vote(ctx context.Context, ld *pr.LowData) (*pr.IsComplited, error) {
	var srr bool
	User = bc.LoadUser(ld.Private, PAREDBNAME)
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_LHASH,
		})
		if res == nil {
			continue
		}
		tx := bc.NewTransaction(User, ld.UserCandidate, bc.Base64Decode(res.Data), ld.Num, ld.Passport, PASSDBNAME)
		res = nt.Send(addr, &nt.Package{
			Option: ADD_TRNSX,
			Data:   bc.SerializeTX(tx),
		})
		if res == nil {
			continue
		}
		if res.Data == "ok" {
			srr = true
		} else {
			srr = false
		}
	}
	return &pr.IsComplited{Ic: srr}, nil
}
func (s *GRserver) ChainPrint(context.Context, *pr.Wpar) (*pr.Chain, error) {
	var allChain []string
	for i := 0; ; i++ {
		res := nt.Send(Addresses[0], &nt.Package{
			Option: GET_BLOCK,
			Data:   fmt.Sprintf("%d", i),
		})
		if res == nil || res.Data == "" {
			break
		}
		fmt.Printf("[%d] => %s\n", i+1, res.Data)
		allChain = append(allChain, fmt.Sprintf("[%d] => %s\n", i+1, res.Data))
	}
	return &pr.Chain{InBlock: allChain}, nil
}*/
