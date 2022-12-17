package FuncLib

import (
	bc "BLCHxVote/Blockchain"
	nt "BLCHxVote/Network"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/valyala/fastjson"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	"io/ioutil"
	"strconv"
	"time"
)

var (
	Addresses []string
	User      *bc.User
	EndTime   string = "8h5m0s"
)

type Proxies []*bc.Candidate
type ListBal []*Horo

type Horo struct {
	Candidate *bc.Candidate
	Balance   string
}

const (
	ADD_BLOCK = iota + 2
	ADD_TRNSX
	GET_BLOCK
	GET_LHASH
	GET_BLNCE
	GET_CSIZE
)
const (
	PASSDBNAME      = "Databases/passdb.db"
	PAREDBNAME      = "Databases/paredb.db"
	PUBLICDBNAME    = "Databases/pubdb.db"
	CANDIDATEDBNAME = "Databases/candidate.db"
)

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

func GenerateDBs() string {
	ChainSize()
	bc.NewVotePass(PASSDBNAME, PAREDBNAME, PUBLICDBNAME, CANDIDATEDBNAME)
	return "ok"
}
func CreatePublic(count int) {
	fmt.Println("1")
	for i := 0; i < count; i++ {
		fmt.Println("2")
		User1 := bc.NewUser(PUBLICDBNAME)
		fmt.Println("3")
		ChainTXBlock(User1.Address(), 1)
		fmt.Println("4")
	}
}
func WinnerList() ListBal {
	var list = ViewCandidates()
	srr := ListBal{}
	db, _ := sql.Open("sqlite3", CANDIDATEDBNAME)
	var desc string
	for i := 0; i < len(list); i++ {
		temp := list[i].PublicKey
		db.QueryRow("SELECT Description FROM CandidateDB WHERE PublicK = $1", temp).Scan(&desc)
		candidate := &bc.Candidate{
			PublicKey:   temp,
			Description: desc,
		}
		p := PrintBalance(candidate.PublicKey)
		srr = append(srr, &Horo{Candidate: candidate, Balance: p})
	}
	return srr
}
func WinnerSolo() *bc.Candidate {
	wl := WinnerList()
	var temp string
	var num1 int
	for i := 0; i < len(wl); i++ {
		if i+1 < len(wl) {
			temp = wl[i+1].Balance
			if wl[i].Balance > temp {
				temp = wl[i].Balance
				num1 = i
			}
			num1 = i + 1
		}
	}
	return wl[num1].Candidate
}
func LimitTime() string {
	temp := ChainBlock("1")
	srr := fastjson.GetString([]byte(temp), "TimeStamp")
	u, _ := time.Parse(time.RFC3339, srr)
	trim := time.Since(u).String()
	return trim
}
func AcceprNewUser(Pass string, PublicK string, salt string) string {
	t, _ := time.ParseDuration(EndTime)
	t1, _ := time.ParseDuration(LimitTime())
	if t1 > t {
		return "2"
	}
	srr := bc.Private(Pass, salt, PASSDBNAME, PAREDBNAME, PublicK, PUBLICDBNAME)
	if srr == "Empty" {
		return "0"
	}
	return srr
}
func AcceprLoadUser(PublicK string, PrivateK string) string {
	t, _ := time.ParseDuration(EndTime)
	t1, _ := time.ParseDuration(LimitTime())
	fmt.Println(t1)
	fmt.Println(t)
	if t1 > t {
		return "2"
	}
	User := bc.LoadUser(PrivateK, PAREDBNAME)
	if User == nil {
		return "0"
	}
	if User.Address() != PublicK {
		return "0"
	}
	return "1"
}
func PrintBalance(useraddr string) string {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	var srr string
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_BLNCE,
			Data:   useraddr,
		})
		if res == nil {
			continue
		}
		srr = fmt.Sprintf("%s", res.Data)
	}
	return srr
}
func ChainSize() string {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	res := nt.Send(Addresses[0], &nt.Package{
		Option: GET_CSIZE,
	})
	if res == nil || res.Data == "" {
		return "fail7"
	}
	srr := fmt.Sprintf("%s", res.Data)
	return srr
}
func ChainPrint() []string {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	var allChain []string
	for i := 0; ; i++ {
		res := nt.Send(Addresses[0], &nt.Package{
			Option: GET_BLOCK,
			Data:   fmt.Sprintf("%d", i),
		})
		if res == nil || res.Data == "" {
			break
		}
		allChain = append(allChain, fmt.Sprintf("%s", res.Data))
	}
	return allChain
}
func ViewCandidates() Proxies {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	db, _ := sql.Open("sqlite3", CANDIDATEDBNAME)
	rows, _ := db.Query("SELECT PublicK, Description FROM CandidateDB")
	p := Proxies{}
	defer db.Close()
	var candidateP string
	var candidateD string
	for rows.Next() {
		rows.Scan(&candidateP, &candidateD)
		p = append(p, &bc.Candidate{PublicKey: candidateP, Description: candidateD})
	}
	return p
}
func ChainTX(candidate string, num uint64, PrivateK string) bool {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	User = bc.LoadUser(PrivateK, PAREDBNAME)
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_LHASH,
		})
		if res == nil {
			continue
		}
		tx := bc.NewTransaction(User, candidate, bc.Base64Decode(res.Data), num)
		res = nt.Send(addr, &nt.Package{
			Option: ADD_TRNSX,
			Data:   bc.SerializeTX(tx),
		})
		if res == nil {
			continue
		}
		if res.Data == "ok" {
			return true
		} else {
			return false
		}
	}
	return false
}
func ChainTXBlock(john string, num uint64) bool {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	for _, addr := range Addresses {
		res := nt.Send(addr, &nt.Package{
			Option: GET_LHASH,
		})
		if res == nil {
			continue
		}
		tx := bc.NewTransactionBlock(john, bc.Base64Decode(res.Data), num)
		res = nt.Send(addr, &nt.Package{
			Option: ADD_TRNSX,
			Data:   bc.SerializeTX(tx),
		})
		if res == nil {
			continue
		}
		if res.Data == "ok" {
			return true
		} else {
			return false
		}
	}

	return false
}
func ChainBlock(splited string) string {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
	num, err := strconv.Atoi(splited)
	if err != nil {
		return "fail3"
	}
	res := nt.Send(Addresses[0], &nt.Package{
		Option: GET_BLOCK,
		Data:   fmt.Sprintf("%d", num-1),
	})
	if res == nil || res.Data == "" {
		return "fail1111"
	}
	srr := fmt.Sprintf("%s", res.Data)
	return srr
}
