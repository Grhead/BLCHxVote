package main

import (
	pr "BLCHxVote/API/Proto"
	"BLCHxVote/Contract"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &Contract.GRserver{}
	pr.RegisterBLCH_ContractServer(s, srv)
	bindPort := ":7070"
	//l, err := net.Listen("tcp4", bindPort)
	l, err := net.Listen("tcp", bindPort)
	if err != nil {
		log.Fatal(err)
	}
	s.Serve(l)
}
