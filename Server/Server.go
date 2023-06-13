package Server

import (
	"VOX2/Basic"
	. "VOX2/Transport/PBs"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

type GRServer struct {
	ContractServer
}

func (s *GRServer) NewChain(ctx context.Context, request *NewChainRequest) (*NewChainResponse, error) {
	chain, err := Basic.NewChain(request.Master, int64(request.VotesCount), request.LimitTime)
	if err != nil {
		return &NewChainResponse{
			CreateHelpProto: nil,
		}, err
	}
	return &NewChainResponse{
		CreateHelpProto: &CreateHelp{Status: chain.Status},
	}, nil
}

func (s *GRServer) CallCreateVoters(ctx context.Context, request *CallCreateVotersRequest) (*CallCreateVotersResponse, error) {
	voters, votersPass, err := Basic.CallCreateVoters(request.Voter, request.Master)
	if err != nil {
		return &CallCreateVotersResponse{User: nil}, err
	}
	var usersList []*BlockchainUser
	for _, v := range voters {
		usersList = append(usersList, &BlockchainUser{
			Id:          v.Id,
			PublicKey:   v.PublicKey,
			IsUsed:      v.IsUsed,
			Affiliation: v.VotingAffiliation,
		})
	}
	return &CallCreateVotersResponse{User: usersList, Identifier: votersPass}, nil
}

func (s *GRServer) CallNewCandidate(ctx context.Context, request *CallNewCandidateRequest) (*CallNewCandidateResponse, error) {
	candidate, err := Basic.CallNewCandidate(request.Description, request.Affiliation)
	log.Println(candidate)
	log.Println(err)
	if err != nil {
		return &CallNewCandidateResponse{ElectionSubjects: nil}, err
	}
	return &CallNewCandidateResponse{ElectionSubjects: &BlockchainElectionSubjects{
		Id:                candidate.Id,
		PublicKey:         candidate.PublicKey,
		Description:       candidate.Description,
		VotingAffiliation: candidate.VotingAffiliation,
	}}, nil
}

func (s *GRServer) CallViewCandidates(ctx context.Context, request *CallViewCandidatesRequest) (*CallViewCandidatesResponse, error) {
	candidates, err := Basic.CallViewCandidates(request.Master)
	if err != nil {
		return &CallViewCandidatesResponse{ElectionSubjects: nil}, err
	}
	var electionsList []*BlockchainElectionSubjects
	for _, v := range candidates {
		electionsList = append(electionsList, &BlockchainElectionSubjects{
			Id:                v.Id,
			PublicKey:         v.PublicKey,
			Description:       v.Description,
			VotingAffiliation: v.VotingAffiliation,
		})
	}
	return &CallViewCandidatesResponse{ElectionSubjects: electionsList}, nil
}

func (s *GRServer) WinnersList(ctx context.Context, request *WinnersListRequest) (*WinnersListResponse, error) {
	list, err := Basic.WinnersList(request.Master)
	if err != nil {
		return &WinnersListResponse{ElectionList: nil}, err
	}
	var resultWinnerList []*ContractElectionsList
	for _, v := range list {
		resultWinnerList = append(resultWinnerList, &ContractElectionsList{
			ElectionSubjects: &BlockchainElectionSubjects{
				Id:                v.ElectionSubject.Id,
				PublicKey:         v.ElectionSubject.PublicKey,
				Description:       v.ElectionSubject.Description,
				VotingAffiliation: v.ElectionSubject.VotingAffiliation,
			},
			Balance: v.Balance,
		})
	}
	return &WinnersListResponse{ElectionList: resultWinnerList}, nil
}

func (s *GRServer) SoloWinner(ctx context.Context, request *SoloWinnerRequest) (*SoloWinnerResponse, error) {
	winner, err := Basic.SoloWinner(request.Master)
	if err != nil {
		return &SoloWinnerResponse{SoloWinnerObject: nil}, err
	}
	return &SoloWinnerResponse{
		SoloWinnerObject: &ContractElectionsList{
			ElectionSubjects: &BlockchainElectionSubjects{
				Id:                winner.ElectionSubject.Id,
				PublicKey:         winner.ElectionSubject.PublicKey,
				Description:       winner.ElectionSubject.Description,
				VotingAffiliation: winner.ElectionSubject.VotingAffiliation,
			},
			Balance: winner.Balance},
	}, nil
}

func (s *GRServer) ChainSize(ctx context.Context, request *ChainSizeRequest) (*ChainSizeResponse, error) {
	size, err := Basic.ChainSize(request.Master)
	if err != nil {
		return &ChainSizeResponse{Size: ""}, err
	}
	return &ChainSizeResponse{Size: size}, nil
}

func (s *GRServer) GetPartOfChain(ctx context.Context, request *GetPartOfChainRequest) (*GetPartOfChainResponse, error) {
	blocks, err := Basic.GetPartOfChain(request.Master)
	if err != nil {
		return &GetPartOfChainResponse{Blocks: nil}, err
	}
	var blocksList []*BlockchainBlock
	for _, v := range blocks {
		var transactionsList []*BlockchainTransaction
		for _, k := range v.Transactions {
			transactionsList = append(transactionsList, &BlockchainTransaction{
				RandBytes: k.RandBytes,
				PrevBlock: k.PrevBlock,
				Sender:    k.Sender,
				Receiver:  k.Receiver,
				Value:     k.Value,
				Signature: k.Signature,
				CurrHash:  k.CurrHash,
			})
		}
		blocksList = append(blocksList, &BlockchainBlock{
			CurrHash:     v.CurrHash,
			PrevHash:     v.PrevHash,
			TimeStamp:    v.TimeStamp,
			Transactions: transactionsList,
			BalanceMap:   v.BalanceMap,
			Nonce:        int64(v.Nonce),
			Difficulty:   v.Difficulty,
			ChainMaster:  v.ChainMaster,
		})
	}
	return &GetPartOfChainResponse{Blocks: blocksList}, nil
}

func (s *GRServer) GetFullChain(ctx context.Context, e *empty.Empty) (*GetFullChainResponse, error) {
	chain, err := Basic.GetFullChain()
	if err != nil {
		return &GetFullChainResponse{Blocks: nil}, err
	}
	var blocksList []*BlockchainBlock
	for _, v := range chain {
		var transactionsList []*BlockchainTransaction
		for _, k := range v.Transactions {
			transactionsList = append(transactionsList, &BlockchainTransaction{
				RandBytes: k.RandBytes,
				PrevBlock: k.PrevBlock,
				Sender:    k.Sender,
				Receiver:  k.Receiver,
				Value:     k.Value,
				Signature: k.Signature,
				CurrHash:  k.CurrHash,
			})
		}
		blocksList = append(blocksList, &BlockchainBlock{
			CurrHash:     v.CurrHash,
			PrevHash:     v.PrevHash,
			TimeStamp:    v.TimeStamp,
			Transactions: transactionsList,
			BalanceMap:   v.BalanceMap,
			Nonce:        int64(v.Nonce),
			Difficulty:   v.Difficulty,
			ChainMaster:  v.ChainMaster,
		})
	}
	return &GetFullChainResponse{Blocks: blocksList}, err
}

func (s *GRServer) AcceptNewUser(ctx context.Context, request *AcceptNewUserRequest) (*AcceptNewUserResponse, error) {
	user, err := Basic.AcceptNewUser(request.Pass, request.Salt, request.PublicKey)
	if err != nil {
		return &AcceptNewUserResponse{PrivateKey: ""}, err
	}
	return &AcceptNewUserResponse{PrivateKey: user}, err
}

func (s *GRServer) AcceptLoadUser(ctx context.Context, request *AcceptLoadUserRequest) (*AcceptLoadUserResponse, error) {
	user, err := Basic.AcceptLoadUser(request.PublicKey, request.PrivateKey)
	if err != nil {
		return &AcceptLoadUserResponse{User: nil}, err
	}
	return &AcceptLoadUserResponse{User: &BlockchainUser{
		Id:          user.Id,
		PublicKey:   user.PublicKey,
		IsUsed:      user.IsUsed,
		Affiliation: user.VotingAffiliation,
	}}, nil
}

func (s *GRServer) Vote(ctx context.Context, request *VoteRequest) (*VoteResponse, error) {
	vote, err := Basic.Vote(request.Receiver, request.Sender, request.Master, request.Num)
	if err != nil {
		return &VoteResponse{Status: ""}, err
	}
	return &VoteResponse{Status: vote}, nil
}

//func (s *GRServer) mustEmbedUnimplementedContractServer() {
//	//TODO implement me
//	panic("implement me")
//}
