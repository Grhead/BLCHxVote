package Network

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"strings"
	"time"
)

const WaitTime = 5
const DMAXSIZE = 2 << 20
const BUFFSIZE = 4096

type Listener net.Listener
type Conn net.Conn
type Package struct {
	Option string
	Data   string
}

var Data string

func init() {
	viper.SetConfigFile("./LowConf/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Send(address string, pack *Package) (*Package, error) {
	fmt.Printf("ooooption: %v\n", pack.Option)
	fmt.Printf("sendsendsend: %v\n", pack.Data)
	EndBytes := viper.GetString("ENDBYTES")
	conn, err := net.Dial("tcp", strings.Trim(address, "\""))
	if err != nil {
		return nil, err
	}
	serializeData, err := SerializePackage(pack)
	if err != nil {
		return nil, err
	}
	_, err = conn.Write([]byte(serializeData + EndBytes))
	if err != nil {
		return nil, err
	}
	res := new(Package)
	ch := make(chan bool)
	go func() {
		res = ReadPackage(conn)
		ch <- true
	}()
	select {
	case <-ch:
	case <-time.After(WaitTime * time.Second):
	}
	return res, nil
}

func ReadPackage(conn net.Conn) *Package {
	fmt.Printf("readdddd: %v\n", conn)
	EndBytes := viper.GetString("ENDBYTES")
	var data string
	var size = uint64(0)
	var buffer = make([]byte, BUFFSIZE)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			log.Fatalln(err)
			return nil
		}
		size += uint64(length)
		if size > DMAXSIZE {
			log.Fatalln(err)
			return nil
		}
		data += string(buffer[:length])
		if strings.Contains(data, EndBytes) {
			data = strings.Split(data, EndBytes)[0]
			break
		}
	}

	return DeserializePackage(data)
}

func handleConn(conn net.Conn, handle func(Conn, *Package)) {
	fmt.Printf("handleCONNNN: %v\n", handle)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	pack := ReadPackage(conn)
	if pack == nil {
		return
	}
	handle(Conn(conn), pack)
}

func serve(listener net.Listener, handle func(Conn, *Package)) {
	/*defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
			return
		}
	}(listener)*/
	defer listener.Close()
	fmt.Printf("start: %v\n", listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go handleConn(conn, handle)
	}
}

func Listen(address string, handle func(Conn, *Package)) Listener {
	splitAddresses := strings.Split(address, ":")
	if len(splitAddresses) != 2 {
		return nil
	}
	listener, err := net.Listen("tcp", "0.0.0.0:"+splitAddresses[1])
	fmt.Printf("Incominge: %v\n", address)
	if err != nil {
		panic(err)
		return nil
	}
	go serve(listener, handle)
	return Listener(listener)
}

func Handle(option string, conn Conn, pack *Package, handle func(*Package) (string, error)) bool {
	fmt.Printf("Incoming package: %v\n", pack)
	EndBytes := viper.GetString("ENDBYTES")
	if pack.Option != option {
		return false
	}
	handledPack, err := handle(pack)
	if err != nil {
		return false
	}
	partOfSerialPack, err := SerializePackage(&Package{
		Option: option,
		Data:   handledPack,
	})
	_, err = conn.Write([]byte(partOfSerialPack + EndBytes))
	if err != nil {
		return false
	}
	return true
}
