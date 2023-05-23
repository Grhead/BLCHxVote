package Basic

import (
	"github.com/valyala/fastjson"
	"os"
)

//func CreateVoters(voter interface{}) (string, error) {
//	switch v := voter.(type) {
//	case int:
//		fmt.Printf("Integer: %v", v)
//	case string:
//		fmt.Printf("String: %v", v)
//	default:
//		fmt.Printf("I don't know, ask stackoverflow.")
//	}
//	return
//}

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
