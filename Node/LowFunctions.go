package main

import (
	"VOX2/Blockchain"
	"VOX2/LowConf"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"sort"
	"strconv"
	"strings"
)

func PushBlockToNet(block *BlockHelp) error {
	type resultStruct struct {
		AddTxStatus string
	}
	var result resultStruct
	var returnErr error
	for _, addr := range OtherAddresses {
		goAddr := addr.String()
		go func() {
			client := req.C().DevMode()
			resp, err := client.R().
				SetBody(&block).
				SetSuccessResult(&result).
				Post(fmt.Sprintf("http://%s/addblock", strings.Trim(goAddr, "\"")))
			fmt.Println(resp.Status)
			if err != nil && !strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
				returnErr = err
				return
			}
		}()
	}
	if returnErr != nil {
		return returnErr
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
	fmt.Println("currSIZE")
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
	Mutex.Unlock()
	if IsMining {
		BreakMining <- true
		IsMining = false
	}

	return "ok", nil
}

func AddTransaction(BlockTx *TransactionHelp) (string, error) {
	if BlockTx.Tx == nil || len(BlockTx.Block.Transactions) == Blockchain.TxsLimit {
		return "", errors.New("transactions limit in blocks")
	}
	Mutex.Lock()
	err := BlockTx.Block.AddTransaction(BlockTx.Tx)
	if err != nil {
		return "", err
	}
	Mutex.Unlock()
	if len(BlockTx.Block.Transactions) == Blockchain.TxsLimit {
		go func() {
			Mutex.Lock()
			goroutineBlock := *BlockTx.Block
			IsMining = true
			Mutex.Unlock()
			res := (&goroutineBlock).Accept(BreakMining)
			Mutex.Lock()
			IsMining = false
			if res == nil && strings.Compare(goroutineBlock.PrevHash, Block.PrevHash) != 0 {
				err = Blockchain.AddBlock(&goroutineBlock)
				if err != nil {
					return
				}
				size, err := Blockchain.Size(goroutineBlock.ChainMaster)
				if err != nil {
					return
				}
				help := BlockHelp{
					Block:   &goroutineBlock,
					Address: ThisServe,
					Size:    size,
				}
				errPBTN := PushBlockToNet(&help)
				if errPBTN != nil {
					return
				}
			}
			Mutex.Unlock()
		}()
	}
	return "ok", nil
}

func CompareChains(address string, master string) error {
	//dbNode, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	//dbCompare, err := gorm.Open(sqlite.Open("Database/CompareDb.db"), &gorm.Config{})
	dbNode, err := gorm.Open(sqlite.Open("DatabaseTest/NodeDb.db"), &gorm.Config{})
	dbCompare, err := gorm.Open(sqlite.Open("DatabaseTest/CompareDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	var blocksResponse []*Blockchain.Block
	client := req.C().DevMode()
	resp, err := client.R().
		SetBody(&master).
		SetSuccessResult(&blocksResponse).
		Post(fmt.Sprintf("http://%s/getblock", strings.Trim(address, "\"")))
	fmt.Println(resp.Status)

	genesis := blocksResponse[0]
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

func NewChain(chainMaster string, count uint64) (string, error) {
	genesis, err := Blockchain.NewChain(count, chainMaster)
	if err != nil {
		return "", nil
	}
	return genesis.CurrHash, nil
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
