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

	l, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	s.Serve(l)
}
