package Contract

import (
	pr "BLCHxVote/API/Proto"
	bc "BLCHxVote/Blockchain"
	cl "BLCHxVote/FuncLib"
	"context"
	"encoding/json"
	"fmt"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	"io/ioutil"
	"strconv"
)

var (
	Addresses []string
	User      *bc.User
)

func init() {
	json.Unmarshal([]byte(readFile("addr.json")), &Addresses)
}
func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

type GRserver struct {
	pr.BLCH_ContractServer
}

func (s *GRserver) mustEmbedUnimplementedBLCH_ContractServer() {
	panic("implement me")
}

func (s *GRserver) AuthRegister(ctx context.Context, ld *pr.RegData) (*pr.AuthRegResult, error) {
	out := cl.AcceprNewUser(ld.Passport, ld.PublicK, ld.Salt)
	return &pr.AuthRegResult{Distortion: out}, nil
}

func (s *GRserver) AuthLogin(ctx context.Context, ld *pr.AuthData) (*pr.AuthRegResult, error) {
	out := cl.AcceprLoadUser(ld.PublicK, ld.PrivateK)
	return &pr.AuthRegResult{Distortion: out}, nil
}

func (s *GRserver) ChainSize(context.Context, *pr.Wpar) (*pr.ResponseSize, error) {
	//srr := cl.ChainSize()
	return &pr.ResponseSize{Size: cl.ChainSize()}, nil
	//return &pr.ResponseSize{Size: cl.LimitTime(time.Now())}, nil
}
func (s *GRserver) TimeBlock(ctx context.Context, ld *pr.BlockDataGet) (*pr.BlockData, error) {
	var srr string
	srr = cl.ChainBlock(ld.BlockNum)
	return &pr.BlockData{InfoBlock: srr, EndTime: cl.EndTime}, nil
}
func (s *GRserver) Balance(ctx context.Context, address *pr.Address) (*pr.Lanb, error) {
	var srr string
	srr = cl.PrintBalance(address.Useradrr)
	return &pr.Lanb{Balance: srr}, nil
}

func (s *GRserver) ResultsWinner(in *pr.Wpar, stream pr.BLCH_Contract_ResultsWinnerServer) error {
	wl := cl.WinnerList()
	var dido float64 = 0.0
	percentInt := 0.0
	for j := 0; j < len(wl); j++ {
		temp, _ := strconv.ParseFloat(wl[j].Balance, 64)
		dido += temp
	}
	for i := 0; i < len(wl); i++ {
		percentInt, _ = strconv.ParseFloat(wl[i].Balance, 64)
		percentString := fmt.Sprintf("%f", (percentInt/dido)*100)
		stream.Send(&pr.CandidateListWithBalance{CandidatePK: wl[i].Candidate.PublicKey, CandidateName: wl[i].Candidate.Description, Balance: percentString})
	}
	return nil
}

func (s *GRserver) ViewCandidates(wr *pr.Wpar, stream pr.BLCH_Contract_ViewCandidatesServer) error {
	results := cl.ViewCandidates()
	for i := 0; i < len(results); i++ {
		temp := results[i]
		stream.Send(&pr.CandidateList{CandidatePK: temp.PublicKey, CandidateName: temp.Description})
	}
	return nil
}
func (s *GRserver) Transfer(ctx context.Context, ld *pr.LowDataChain) (*pr.IsComplited, error) {
	var srr bool
	srr = cl.ChainTXBlock(ld.UserCandidate, ld.Num)
	return &pr.IsComplited{Ic: srr}, nil
}
func (s *GRserver) Vote(ctx context.Context, ld *pr.LowData) (*pr.IsComplited, error) {
	var srr bool
	srr = cl.ChainTX(ld.UserCandidate, ld.Num, ld.Private)
	return &pr.IsComplited{Ic: srr}, nil
}
func (s *GRserver) ChainPrint(context.Context, *pr.Wpar) (*pr.Chain, error) {
	var allChain []string
	allChain = cl.ChainPrint()
	return &pr.Chain{InBlock: allChain}, nil
}
