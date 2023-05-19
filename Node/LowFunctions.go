package main

import (
	"VOX2/Blockchain"
	"VOX2/LowConf"
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func NewChain(chainMaster string, count uint64) (string, error) {
	genesis, err := Blockchain.NewChain(count, chainMaster)
	if err != nil {
		return "", nil
	}
	return genesis.CurrHash, nil
}

func PushBlockToNet(block *Blockchain.Block) error {
	serialBlock, err := Blockchain.SerializeBlock(block)
	if err != nil {
		return err
	}
	//chainSizeForMsg, err := Blockchain.Size(block.ChainMaster)
	//if err != nil {
	//	return err
	//}
	//var msg = ThisServe +
	//	LowConf.Separator +
	//	fmt.Sprintf("%s", block.ChainMaster) +
	//	LowConf.Separator +
	//	fmt.Sprintf("%d", chainSizeForMsg) +
	//	LowConf.Separator +
	//	serialBlock
	for _, addr := range OtherAddresses {
		goAddr := addr.String()
		go func() {
			resp, err := http.Post(fmt.Sprintf("https:/%s/addblock", goAddr),
				"application/json",
				bytes.NewBuffer([]byte(serialBlock)))
			if err != nil {
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					return
				}
			}(resp.Body)
		}()
		//_, err := Network.Send(goAddr, &Network.Package{
		//	Option: LowConf.AddBlockConst,
		//	Data:   msg,
		//})
	}
	return nil
}

func AddBlock(pack *BlockHelp) (string, error) {
	block := pack.Block
	currSize, err := Blockchain.Size(block.ChainMaster)
	if err != nil {
		return "", err
	}
	num := pack.Size
	if currSize < num {
		go func() {
			err := CompareChains(pack.Address, block.ChainMaster)
			if err != nil {
				return
			}
		}()
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

func AddTransaction(tx *Blockchain.Transaction) (string, error) {
	//tx, err := Blockchain.DeserializeTX(pack.Data)
	//if err != nil {
	//	return "", err
	//}
	if tx == nil || len(Block.Transactions) == Blockchain.TxsLimit {
		return "", errors.New("transactions limit in blocks")
	}
	Mutex.Lock()
	err := Block.AddTransaction(tx)
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
			res := (&block).Accept(BreakMining)
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

func CompareChains(address string, master string) error {
	dbNode, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	dbCompare, err := gorm.Open(sqlite.Open("Database/CompareDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	res0, err := Network.Send(address, &Network.Package{
		Option: LowConf.GetBlockConst,
		//Data:   fmt.Sprintf("%d", 0),
		Data: fmt.Sprintf("%s", master),
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
	//TODO ERROR
	for i := 1; i < 10; i++ {
		res1, err := Network.Send(address, &Network.Package{
			Option: LowConf.GetBlockConst,
			//Data:   fmt.Sprintf("%d", i),
			Data: fmt.Sprintf("%s", i),
		})
		if err != nil {
			return err
		}
		if res1 == nil {
			return errors.New("request data is nil")
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

func GetBlocks(pack *MasterHelp) ([]*Blockchain.Block, error) {
	blocks, err := Blockchain.GetFullChain(pack.Master)
	if err != nil {
		return nil, err
	}
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].TimeStamp.AsTime().After(blocks[j].TimeStamp.AsTime())
	})
	return blocks, nil
}

func GetLastHash(pack *MasterHelp) (string, error) {
	return Blockchain.LastHash(pack.Master)
}
func GetBalance(pack *UserHelp) (string, error) {
	log.Println("Get-Balance")
	balance, err := Blockchain.Balance(pack.User)
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(balance, 10), nil
}
func GetChainSize(pack *MasterHelp) (string, error) {
	size, err := Blockchain.Size(pack.Master)
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(size, 10), nil
}
