package main

import (
	"VOX2/Server"
	. "VOX2/Transport/PBs"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &Server.GRServer{}
	RegisterContractServer(s, srv)
	bindPort := ":9099"
	l, err := net.Listen("tcp", bindPort)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Serve(l)
	if err != nil {
		return
	}
}
