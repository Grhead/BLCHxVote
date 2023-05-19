package Basic

import (
	"VOX2/LowConf"
	"VOX2/Transport/Network"
	"fmt"
	"github.com/valyala/fastjson"
	"os"
)

func PrintBalance(moneyMan string) (string, error) {
	addresses, err := readAddresses()
	if err != nil {
		return "", err
	}
	var manBalance string
	for _, address := range addresses {
		response, errSend := Network.Send(address.String(), &Network.Package{
			Option: LowConf.GetBalanceConst,
			Data:   moneyMan,
		})
		if errSend != nil {
			return "", errSend
		}
		if response == nil {
			continue
		}
		manBalance = fmt.Sprintf("%s", response.Data)
	}
	return manBalance, nil
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
