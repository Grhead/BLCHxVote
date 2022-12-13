package main

import (
	"io/ioutil"
)

var (
	Addresses []string
)

const (
	SEPARATOR = "_SEPARATOR_"
	ADD_BLOCK = iota + 1
	ADD_TRNSX
	GET_BLOCK
	GET_LHASH
	GET_BLNCE
	GET_CSIZE
)

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}
