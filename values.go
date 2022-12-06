package main

import (
	bc "BLCHxVote/Blockchain"
	"io/ioutil"
)

var (
	Addresses []string
	User      *bc.User
)

const (
	SEPARATOR = "_SEPARATOR_"
	ADD_BLOCK = iota + 1
	ADD_TRNSX
	GET_BLOCK
	GET_LHASH
	GET_BLNCE
	GET_CSIZE
	DBNAME       = "Databases/bloch.db"
	PASSDBNAME   = "Databases/passdb.db"
	PAREDBNAME   = "Databases/paredb.db"
	PUBLICDBNAME = "Databases/pubdb.db"
)

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}
