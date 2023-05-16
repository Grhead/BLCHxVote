package Node

import (
	"VOX2/Blockchain"
	"VOX2/Transport/Network"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"sync"
)

var Mutex sync.Mutex
var IsMining bool
var BreakMining = make(chan bool)
var Block *Blockchain.Block

// TODO CreateNew			// TODO NewChain
// TODO CompareChains		// TODO NewBlock
// TODO PushBlockToNet		// TODO NewTransaction
// TODO AddBlock			// TODO NewTransactionFromChain
// TODO AddTransaction		// TODO LastHash
// TODO GetBlock			// TODO AddBlock
// TODO GetLastHash			// TODO NewDormantUser
// TODO GetBalance			// TODO LoadToEnterAlreadyUser
// TODO GetChainSize		// TODO NewPublicKeyItem
// TODO SelectBlock			// TODO NewCandidate
// TODO HashBlock			// TODO Size
// TODO CopyFile			// TODO Balance

// TODO RegisterGeneratePrivate
// TODO GenerateKey

func CompareChains(address string, chainSize int) error {
	dbNode, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	dbCompare, err := gorm.Open(sqlite.Open("Database/CompareDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	res0, err := Network.Send(address, &Network.Package{
		Option: GetBlock,
		Data:   fmt.Sprintf("%d", 0),
	})
	if err != nil {
		return err
	}
	genesis, err := Blockchain.DeserializeBlock(res0.Data)
	if err != nil {
		return err
	}
	if strings.Compare(genesis.CurrHash, genesis.Hash()) != 0 {
		return errors.New("hashes are not the same")
	}
	err = Blockchain.AddBlockCompare(genesis)
	if err != nil {
		return err
	}
	for i := 1; i < chainSize; i++ {
		res1, err := Network.Send(address, &Network.Package{
			Option: GetBlock,
			Data:   fmt.Sprintf("%d", i),
		})
		if err != nil {
			return err
		}
		if res1 == nil {
			return errors.New("block is nil")
		}
		block, errDeserialize := Blockchain.DeserializeBlock(res1.Data)
		if errDeserialize != nil {
			return errDeserialize
		}
		if block == nil {
			return errors.New("block is nil")
		}
		errAddBlock := Blockchain.AddBlockCompare(block)
		if errAddBlock != nil {
			return errAddBlock
		}
	}
	Mutex.Lock()
	var blocks []*Blockchain.Chain
	dbCompare.Find(&blocks)
	errDelete := dbNode.Exec("DELETE FROM Chains")
	if errDelete.Error != nil {
		return errDelete.Error
	}
	for _, v := range blocks {
		errInsert := dbNode.Exec("INSERT INTO Chains (Id, Hash, Block) VALUES ($1, $2, $3)",
			uuid.NewString(),
			v.Hash,
			v.Block,
		)
		if errInsert.Error != nil {
			return errInsert.Error
		}
	}
	errDelete = dbCompare.Exec("DELETE FROM Chains")
	if errDelete.Error != nil {
		return errDelete.Error
	}
	lastHash, err := Blockchain.LastHash(Block.ChainMaster)
	if err != nil {
		return err
	}
	Block, err = Blockchain.NewBlock(Block.CurrHash, lastHash)
	if err != nil {
		return err
	}
	Mutex.Unlock()
	if IsMining {
		BreakMining <- true
		IsMining = false
	}
	return nil
}
