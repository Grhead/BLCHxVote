package main

import (
	"VOX2/Server"
	. "VOX2/Transport/PBs"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	viper.SetConfigFile("./LowConf/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	srv := &Server.GRServer{}
	RegisterContractServer(s, srv)
	bindPort := viper.GetString("BIND_PORT")
	l, err := net.Listen("tcp", bindPort)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Serve(l)
	if err != nil {
		return
	}
}
