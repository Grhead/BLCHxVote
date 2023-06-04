package Server

import (
	. "VOX2/Transport/PBs"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type GRServer struct {
	ContractServer
}

func (s *GRServer) NewChain(ctx context.Context, request *NewChainRequest) (*NewChainResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) CallCreateVoters(ctx context.Context, request *CallCreateVotersRequest) (*CallCreateVotersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) CallNewCandidate(ctx context.Context, request *CallNewCandidateRequest) (*CallNewCandidateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) CallViewCandidates(ctx context.Context, request *CallNewCandidateRequest) (*CallNewCandidateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) WinnersList(ctx context.Context, request *WinnersListRequest) (*WinnersListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) SoloWinner(ctx context.Context, request *SoloWinnerRequest) (*SoloWinnerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) ChainSize(ctx context.Context, request *ChainSizeRequest) (*ChainSizeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) GetPartOfChain(ctx context.Context, request *GetPartOfChainRequest) (*GetPartOfChainResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) GetFullChain(ctx context.Context, e *empty.Empty) (*GetFullChainResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) AcceptNewUser(ctx context.Context, request *AcceptNewUserRequest) (*AcceptNewUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) AcceptLoadUser(ctx context.Context, request *AcceptLoadUserRequest) (*AcceptLoadUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) Vote(ctx context.Context, request *VoteRequest) (*VoteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRServer) mustEmbedUnimplementedContractServer() {
	//TODO implement me
	panic("implement me")
}
