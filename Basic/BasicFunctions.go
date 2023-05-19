package Basic

import (
	"VOX2/LowConf"
	"VOX2/Transport/Network"
	"fmt"
	"github.com/valyala/fastjson"
	"os"
	"strings"
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
		if response == nil {
			continue
		}
		if errSend != nil && !strings.Contains(errSend.Error(), "No connection could be made because the target machine actively refused it.") {
			return "", errSend
		}
		manBalance = fmt.Sprintf("%v", response.Data)
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
