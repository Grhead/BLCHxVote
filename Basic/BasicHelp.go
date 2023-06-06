package Basic

import (
	"VOX2/Transport"
	//"encoding/json"
	"github.com/goccy/go-json"
	"github.com/valyala/fastjson"
	"os"
)

func ReadAddresses() ([]*fastjson.Value, error) {
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

func SerializeTX(tx *Transport.TransactionHelp) (string, error) {
	jsonData, err := json.MarshalIndent(*tx, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func DeserializeTX(data *string) *Transport.TransactionHelp {
	var tx Transport.TransactionHelp
	err := json.Unmarshal([]byte(*data), &tx)
	if err != nil {
		return nil
	}
	return &tx
}
