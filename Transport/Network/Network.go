package Network

import (
	"encoding/json"
	"net"
	"strings"
	"time"
)

const (
	ENDBYTES = "\000\001\000\001\000"
	WAITTIME = 5
	DMAXSIZE = 2 << 20
	BUFFSIZE = 4096
)

type Package struct {
	Option int
	Data   string
}

type Listener net.Listener
type Conn net.Conn

func Send(address string, pack *Package) *Package {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil
	}
	serializeData, err := SerializePackage(pack)
	if err != nil {
		return nil
	}
	_, err = conn.Write([]byte(serializeData + ENDBYTES))
	if err != nil {
		return nil
	}
	var res = new(Package)

	ch := make(chan bool)
	go func() {
		res = ReadPackage(conn)
		ch <- true
	}()
	select {
	case <-ch:
	case <-time.After(WAITTIME * time.Second):
	}
	return res
}

func ReadPackage(conn net.Conn) *Package {
	var data string
	var size = uint64(0)
	var buffer = make([]byte, BUFFSIZE)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			return nil
		}
		size += uint64(length)
		if size > DMAXSIZE {
			return nil
		}
		data += string(buffer[:length])
		if strings.Contains(data, ENDBYTES) {
			data = strings.Split(data, ENDBYTES)[0]
			break
		}
	}

	return DeserializePackage(data)
}

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
