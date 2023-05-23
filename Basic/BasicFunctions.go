package Basic

import (
	"VOX2/Blockchain"
	"errors"
	"fmt"
	"github.com/valyala/fastjson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strconv"
)

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
