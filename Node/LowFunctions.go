package main

import (
	"VOX2/Blockchain"
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

type resultStruct struct {
	AddTxStatus string
}

func goAddBlock(block *BlockHelp, result resultStruct, goAddr string) {
	client := req.C().DevMode()
	_, err := client.R().
		SetBody(&block).
		SetSuccessResult(&result).
		Post(fmt.Sprintf("http://%s/addblock", strings.Trim(goAddr, "\"")))
	if err != nil && !strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
		log.Fatal(err)
	}
}
func goAddTransaction(BlockTx *TransactionHelp) {
	Mutex.Lock()
	goroutineBlock := *BlockTx.Block
	IsMining = true
	Mutex.Unlock()
	res := (&goroutineBlock).Accept(BreakMining)
	Mutex.Lock()
	IsMining = false
	if res == nil && strings.Compare(goroutineBlock.PrevHash, Block.PrevHash) != 0 {
		err := Blockchain.AddBlock(&goroutineBlock)
		if err != nil {
			log.Fatal(err)
		}
		size, err := Blockchain.Size(goroutineBlock.ChainMaster)
		if err != nil {
			log.Fatal(err)
		}
		help := BlockHelp{
			Block:   &goroutineBlock,
			Address: ThisServe,
			Size:    size,
		}
		err = PushBlockToNet(&help)
		if err != nil {
			log.Fatal(err)
		}
	}
	Mutex.Unlock()
}
func goCompare(address string, chainMaster string) {
	err := CompareChains(address, chainMaster)
	if err != nil {
		log.Fatal(err)
	}
}

func PushBlockToNet(block *BlockHelp) error {
	var result resultStruct
	var returnErr error
	for _, addr := range OtherAddresses {
		goAddr := addr.String()
		goAddBlock(block, result, goAddr)
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
	if currSize < num {
		fmt.Println("inside")
		goCompare(pack.Address, block.ChainMaster)
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
		goAddTransaction(BlockTx)
	}
	return "ok", nil
}

func CompareChains(address string, master string) error {
	dbNode, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	dbCompare, err := gorm.Open(sqlite.Open("Database/CompareDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	var blocksResponse []*Blockchain.Block
	var masterSend MasterHelp
	masterSend.Master = master
	client := req.C().DevMode()
	_, err = client.R().
		SetBody(&masterSend).
		SetSuccessResult(&blocksResponse).
		Post(fmt.Sprintf("http://%s/getblock", strings.Trim(address, "\"")))
	if err != nil {
		return err
	}
	fmt.Println("================================================================================")
	genesis := blocksResponse[0]
	if strings.Compare(genesis.CurrHash, genesis.Hash()) != 0 {
		return errors.New("hashes are not the same")
	}
	err = Blockchain.AddBlockCompare(genesis)
	if err != nil {
		return err
	}
	size, err := Blockchain.Size(master)
	if err != nil {
		return err
	}
	//TODO ERROR
	for i := uint64(1); i < size; i++ {
		block := blocksResponse[i]
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
	//errDelete := dbNode.Exec("DELETE FROM Chains")
	//if errDelete.Error != nil {
	//	return errDelete.Error
	//}
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
	//errDelete = dbCompare.Exec("DELETE FROM Chains")
	//if errDelete.Error != nil {
	//	return errDelete.Error
	//}
	//lastHash, err := Blockchain.LastHash(Block.ChainMaster)
	//if err != nil {
	//	return err
	//}
	//Block, err = Blockchain.NewBlock(Block.CurrHash, lastHash)
	//if err != nil {
	//	return err
	//}
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
