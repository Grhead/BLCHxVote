package Basic

import (
	"VOX2/Blockchain"
	"errors"
	"fmt"
	"github.com/valyala/fastjson"
	"os"
)

func CreateVoters(voter interface{}, master string) (string, error) {
	addresses, err := readAddresses()
	if err != nil {
		return "", err
	}
	switch v := voter.(type) {
	case int:
		item, err := Blockchain.NewPublicKeyItem(master)
		if err != nil {
			return "", err
		}
	case string:
		err := Blockchain.NewDormantUser(fmt.Sprintf("%v", voter))
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("invalid type")
	}
	return "ok", nil
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

//client := req.C().DevMode()
//_, err := client.R().
//SetBody(&block).
//SetSuccessResult(&result).
//Post(fmt.Sprintf("http://%s/addblock", strings.Trim(goAddr, "\"")))
