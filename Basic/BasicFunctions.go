package Basic

import (
	"VOX2/Blockchain"
	"VOX2/Transport"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/valyala/fastjson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

type UserHelp struct {
	User string `form:"user" json:"user"`
}
type BalanceHelp struct {
	Balance string `form:"balance" json:"balance"`
}
type MasterHelp struct {
	Master string `form:"master" json:"master"`
}
type SizeHelp struct {
	ChainSize string `form:"chainSize" json:"chainSize"`
}

func CallCreateVoters(voter interface{}, master string) ([]*Blockchain.User, error) {
	var resultItems []*Blockchain.User
	switch voter.(type) {
	case int:
		max, err := strconv.Atoi(fmt.Sprintf("%v", voter))
		if err != nil {
			return nil, err
		}
		for i := 0; i < max; i++ {
			item, errNewPublicKey := Blockchain.NewPublicKeyItem(master)
			if errNewPublicKey != nil {
				return nil, errNewPublicKey
			}
			resultItems = append(resultItems, item)
		}
	case string:
		err := Blockchain.NewDormantUser(fmt.Sprintf("%v", voter))
		if err != nil {
			return nil, err
		}
		item, err := Blockchain.NewPublicKeyItem(master)
		if err != nil {
			return nil, err
		}
		resultItems = append(resultItems, item)
	default:
		return nil, errors.New("invalid type")
	}
	return resultItems, nil
}

func CallViewCandidates() ([]*Blockchain.ElectionSubjects, error) {
	db, err := gorm.Open(sqlite.Open("Database/ContractDB.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var Candidates []*Blockchain.ElectionSubjects
	db.Table("ElectionSubjects").Find(&Candidates)
	return Candidates, nil
}

func CallNewCandidate(description string, affiliation string) (*Blockchain.ElectionSubjects, error) {
	candidate, err := Blockchain.NewCandidate(description, affiliation)
	if err != nil {
		return nil, err
	}
	return candidate, nil
}

func GetBalance(userAddress string) (string, error) {
	addresses, err := readAddresses()
	if err != nil {
		return "", err
	}
	userAddressStruct := UserHelp{
		User: userAddress,
	}
	var userBalance *BalanceHelp
	client := req.C().DevMode()
	for _, addr := range addresses {
		resp, errReq := client.R().
			SetBody(&userAddressStruct).
			SetSuccessResult(&userBalance).
			Post(fmt.Sprintf("http://%s/getbalance", strings.Trim(addr.String(), "\"")))
		if errReq != nil {
			return "", errReq
		}
		if resp.Body == nil {
			continue
		}
	}
	return userBalance.Balance, nil
}

func ChainSize(master string) (string, error) {
	addresses, err := readAddresses()
	if err != nil {
		return "", err
	}
	masterChain := MasterHelp{Master: master}
	var chainSize SizeHelp
	client := req.C().DevMode()
	resp, err := client.R().
		SetBody(&masterChain).
		SetSuccessResult(&chainSize).
		Post(fmt.Sprintf("http://%s/getchainsize", strings.Trim(addresses[0].String(), "\"")))
	if err != nil {
		return "", err
	}
	if resp.Body == nil {
		return "", errors.New("empty response")
	}
	return chainSize.ChainSize, nil
}

func GetPartOfChain(master string) ([]*Blockchain.Block, error) {
	addresses, err := readAddresses()
	if err != nil {
		return nil, err
	}
	var partOfChain []*Blockchain.Block
	ChainMaster := Transport.MasterHelp{
		Master: master,
	}
	client := req.C().DevMode()
	resp, err := client.R().
		SetBody(&ChainMaster).
		SetSuccessResult(&partOfChain).
		Post(fmt.Sprintf("http://%s/getblock", strings.Trim(addresses[0].String(), "\"")))
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, errors.New("empty response")
	}
	return partOfChain, nil
}

func GetFullChain() ([]*Blockchain.Block, error) {
	addresses, err := readAddresses()
	if err != nil {
		return nil, err
	}
	var fullChain []*Blockchain.Block
	client := req.C().DevMode()
	resp, err := client.R().SetSuccessResult(&fullChain).
		Get(fmt.Sprintf("http://%s/getdb", strings.Trim(addresses[0].String(), "\"")))
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, errors.New("empty response")
	}
	return fullChain, nil
}

func AcceptNewUser(Pass string, salt string, PublicKey string) (string, error) {
	//t, _ := time.ParseDuration(EndTime)
	//t1, _ := time.ParseDuration(LimitTime())
	//if t1 > t {
	//	return "time"
	//}
	private, err := Blockchain.RegisterGeneratePrivate(Pass, salt, PublicKey)
	if err != nil {
		return "", err
	}
	return private, nil
}

func readAddresses() ([]*fastjson.Value, error) {
	file, err := os.ReadFile("LowConf/addr.json")
	if err != nil {
		return nil, err
	}
	var p fastjson.Parser
	v, err := p.Parse(string(file))
	if err != nil {
		return nil, err
	}
	return v.GetArray("addresses"), nil
}

//addresses, err := readAddresses()
//if err != nil {
//return "", err
//}

//client := req.C().DevMode()
//_, err := client.R().
//SetBody(&block).
//SetSuccessResult(&result).
//Post(fmt.Sprintf("http://%s/addblock", strings.Trim(goAddr, "\"")))
