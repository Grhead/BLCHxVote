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

var BlockForTransaction *Blockchain.Block

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
func goAddTransaction() {
	Mutex.Lock()
	goroutineBlock := *BlockForTransaction
	IsMining = true
	Mutex.Unlock()
	res := (&goroutineBlock).Accept(BreakMining)
	Mutex.Lock()
	IsMining = false
	if res == nil && strings.Compare(goroutineBlock.PrevHash, BlockForTransaction.PrevHash) != 0 {
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
		err = pushBlockToNet(&help)
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

func pushBlockToNet(block *BlockHelp) error {
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
	fmt.Println("out")
	err = Blockchain.AddBlock(block)
	if err != nil {
		return "", err
	}
	Mutex.Unlock()
	if IsMining {
		BreakMining <- true
		IsMining = false
	}
	BlockForTransaction = nil
	return "ok", nil
}

func AddTransaction(BlockTx *TransactionHelp) (string, error) {
	if BlockTx.Tx == nil || len(BlockForTransaction.Transactions) == Blockchain.TxsLimit {
		return "", errors.New("transactions limit in blocks")
	}
	hash, err := Blockchain.LastHash(BlockTx.Master)
	if err != nil {
		return "", err
	}
	BlockForTransaction, err = Blockchain.NewBlock(hash, BlockTx.Master)
	if err != nil {
		return "", err
	}
	Mutex.Lock()
	err = BlockForTransaction.AddTransaction(BlockTx.Tx)
	if err != nil {
		return "", err
	}
	Mutex.Unlock()
	if len(BlockForTransaction.Transactions) == Blockchain.TxsLimit {
		goAddTransaction()
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
	client := req.C().DevMode()
	_, err = client.R().SetSuccessResult(&blocksResponse).
		Get(fmt.Sprintf("http://%s/getdb", strings.Trim(address, "\"")))
	if err != nil {
		return err
	}
	size, err := Blockchain.DbSize()
	if err != nil {
		return err
	}

	someGenesis := blocksResponse[0]
	if strings.Compare(someGenesis.CurrHash, someGenesis.Hash()) != 0 {
		return errors.New("hashes are not the same")
	}
	err = Blockchain.AddBlockCompare(someGenesis)
	if err != nil {
		return err
	}
	for i := uint64(1); i < size; i++ {
		fmt.Println("--------", i)
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
	BlockForTransaction = nil
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
