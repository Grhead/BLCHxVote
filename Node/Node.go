package Node

import (
	"VOX2/Blockchain"
	"VOX2/Transport/Network"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/valyala/fastjson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var Mutex sync.Mutex
var IsMining bool
var BreakMining = make(chan bool)
var Block *Blockchain.Block
var ThisServe string
var OtherAddresses []*fastjson.Value

// TODO --CreateNew			// TODO NewChain
// TODO --CompareChains		// TODO NewBlock
// TODO --PushBlockToNet		// TODO NewTransaction
// TODO --AddBlock			// TODO NewTransactionFromChain
// TODO --AddTransaction		// TODO LastHash
// TODO --GetBlock_const			// TODO AddBlock
// TODO GetLastHash			// TODO NewDormantUser
// TODO GetBalance			// TODO LoadToEnterAlreadyUser
// TODO GetChainSize		// TODO NewPublicKeyItem
// TODO SelectBlock			// TODO NewCandidate
// TODO HashBlock			// TODO Size
// TODO CopyFile			// TODO Balance

// TODO RegisterGeneratePrivate
// TODO GenerateKey

func init() {
	ThisServe = ":7575"
	file, err := os.ReadFile("LowConf/addr.json")
	if err != nil {
		log.Fatalln(err)
	}
	var p fastjson.Parser
	v, err := p.Parse(string(file))
	if err != nil {
		log.Fatalln(err)
	}
	OtherAddresses = v.GetArray("addresses")
	fmt.Println(OtherAddresses)
}

func NewChain(chainMaster string, count uint64) error {
	_, err := Blockchain.NewChain(count, chainMaster)
	if err != nil {
		return nil
	}
	return err
}

func PushBlockToNet(block *Blockchain.Block) error {
	sblock, err := Blockchain.SerializeBlock(block)
	if err != nil {
		return err
	}
	chainSizeForMsg, err := Blockchain.Size(block.ChainMaster)
	if err != nil {
		return err
	}
	var msg = ThisServe + Separator + fmt.Sprintf("%d", chainSizeForMsg) + Separator + sblock
	for _, addr := range OtherAddresses {
		goAddr := addr.String()
		go func() {
			_, err := Network.Send(goAddr, &Network.Package{
				Option: AddBlock_const,
				Data:   msg,
			})
			if err != nil {
				return
			}
		}()
	}
	return nil
}

func AddBlock(pack *Network.Package) (string, error) {
	splited := strings.Split(pack.Data, Separator)
	block, err := Blockchain.DeserializeBlock(splited[2])
	if err != nil {
		return "", err
	}
	currSize, err := Blockchain.Size(block.ChainMaster)
	if err != nil {
		return "", err
	}
	num, _ := strconv.Atoi(splited[1])
	if currSize < uint64(num) {
		err := CompareChains(splited[0], num)
		if err != nil {
			return "", err
		}
		return "ok ", nil
	}

	Mutex.Lock()
	err = Blockchain.AddBlock(block)
	if err != nil {
		return "", err
	}
	lastHash, err := Blockchain.LastHash(block.ChainMaster)
	if err != nil {
		return "", err
	}
	Block, err = Blockchain.NewBlock(block.CurrHash, lastHash)
	if err != nil {
		return "", err
	}
	Mutex.Unlock()
	if IsMining {
		BreakMining <- true
		IsMining = false
	}

	return "ok", nil
}

func addTransaction(pack *Network.Package) (string, error) {
	tx, err := Blockchain.DeserializeTX(pack.Data)
	if err != nil {
		return "", err
	}
	if tx == nil || len(Block.Transactions) == Blockchain.TxsLimit {
		return "", errors.New("transactions limit in blocks")
	}
	Mutex.Lock()
	err = Block.AddTransaction(tx, Block.ChainMaster)
	if err != nil {
		return "", err
	}
	Mutex.Unlock()
	if len(Block.Transactions) == Blockchain.TxsLimit {
		go func() {
			Mutex.Lock()
			block := *Block
			IsMining = true
			Mutex.Unlock()
			user, err := Blockchain.FindByEnterUserWithLogin(tx.Sender)
			if err != nil {
				return
			}
			res := (&block).Accept(user, Block.ChainMaster, BreakMining)
			Mutex.Lock()
			IsMining = false
			if res == nil && strings.Compare(block.PrevHash, Block.PrevHash) != 0 {
				err = Blockchain.AddBlock(&block)
				if err != nil {
					return
				}
				err := PushBlockToNet(&block)
				if err != nil {
					return
				}
			}
			lastHash, err := Blockchain.LastHash(Block.ChainMaster)
			if err != nil {
				return
			}
			Block, err = Blockchain.NewBlock(Block.CurrHash, lastHash)
			if err != nil {
				return
			}
			Mutex.Unlock()
		}()
	}
	return "ok", nil
}

func CompareChains(address string, chainSize int) error {
	dbNode, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	dbCompare, err := gorm.Open(sqlite.Open("Database/CompareDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	res0, err := Network.Send(address, &Network.Package{
		Option: GetBlock_const,
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
			Option: GetBlock_const,
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
