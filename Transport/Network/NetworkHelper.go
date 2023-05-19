package Network

import "encoding/json"

func SerializePackage(pack *Package) (string, error) {
	jsonData, err := json.MarshalIndent(*pack, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func DeserializePackage(data string) *Package {
	var pack Package
	err := json.Unmarshal([]byte(data), &pack)
	if err != nil {
		return nil
	}
	return &pack
}
