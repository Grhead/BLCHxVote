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
func goAddTransaction() {
	Mutex.Lock()
	goroutineBlock := BlockForTransaction
	IsMining = true
	Mutex.Unlock()

	/*if IsMining &&  {

	}*/

	res := (goroutineBlock).Accept(BreakMining)
	Mutex.Lock()
	IsMining = false
	if res == nil && strings.Compare(goroutineBlock.PrevHash, BlockForTransaction.PrevHash) == 0 {
		fmt.Println("=====", 1)
		err := Blockchain.AddBlock(goroutineBlock)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("=====", 2)
		size, err := Blockchain.Size(goroutineBlock.ChainMaster)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("=====", 3)
		help := BlockHelp{
			Block:   goroutineBlock,
			Address: ThisServe,
			Size:    size,
		}
		err = pushBlockToNet(&help)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("=====", 4)
	}
	fmt.Println("=====NOOOOOOOOOOOOOOOOOOOOO")
	hash, err := Blockchain.LastHash("Start")
	if err != nil {
		log.Fatal(err)
	}
	BlockForTransaction, err = Blockchain.NewBlock(hash, "Start")
	if err != nil {
		log.Fatal(err)
	}
	Mutex.Unlock()
}
func goCompare(address string) {
	err := CompareChains(address)
	if err != nil {
		log.Fatal(err)
	}
}

func pushBlockToNet(block *BlockHelp) error {
	//err := Blockchain.AddBlock(block.Block)
	//if err != nil {
	//	log.Fatal(err)
	//}
	var result resultStruct
	var returnErr error
	for _, addr := range OtherAddresses {
		goAddr := addr.String()
		go goAddBlock(block, result, goAddr)
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
		go goCompare(pack.Address)
		return "ok", nil
	}
	Mutex.Lock()
	fmt.Println("out")
	//err = Blockchain.AddBlock(block)
	//if err != nil {
	//	return "", err
	//}
	Mutex.Unlock()
	if IsMining {
		BreakMining <- true
		IsMining = false
	}
	hash, err := Blockchain.LastHash("Start")
	if err != nil {
		return "", nil
	}
	BlockForTransaction, err = Blockchain.NewBlock(hash, "Start")
	if err != nil {
		return "", nil
	}
	return "ok", nil
}

func AddTransaction(BlockTx *TransactionHelp) (string, error) {
	if BlockTx.Tx == nil {
		return "", errors.New("tx is empty")
	}
	if len(BlockForTransaction.Transactions) == Blockchain.TxsLimit {
		WaitTransaction = append(WaitTransaction, BlockTx.Tx)
		//return "", errors.New("transactions limit in blocks")
	}
	Mutex.Lock()
	err := BlockForTransaction.AddTransaction(BlockTx.Tx)
	if err != nil {
		return "", err
	}
	Mutex.Unlock()
	//fmt.Println(len(BlockForTransaction.Transactions))

	if len(BlockForTransaction.Transactions) == Blockchain.TxsLimit {
		go goAddTransaction()
	}
	if len(WaitTransaction) != 0 ** {
		Mutex.Lock()
		err := BlockForTransaction.AddTransaction(BlockTx.Tx)
		if err != nil {
			return "", err
		}
		Mutex.Unlock()
		if len(BlockForTransaction.Transactions) == Blockchain.TxsLimit {
			go goAddTransaction()
		}
	}
	return "ok", nil
}

func CompareChains(address string) error {
	dbNode, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
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
	someGenesis := blocksResponse[0]
	if strings.Compare(someGenesis.CurrHash, someGenesis.Hash()) != 0 {
		return errors.New("hashes are not the same")
	}
	var arrayToMerge []*Blockchain.Chain
	serializeGenesisBlock, err := Blockchain.SerializeBlock(someGenesis)
	fmt.Println("--------00")
	if err != nil {
		return err
	}
	genesisBlock := Blockchain.Chain{
		Id:    uuid.NewString(),
		Hash:  someGenesis.CurrHash,
		Block: serializeGenesisBlock,
	}
	arrayToMerge = append(arrayToMerge, &genesisBlock)
	for _, v := range blocksResponse {
		if v != blocksResponse[0] {
			fmt.Println("--------", v)
			if v == nil {
				return errors.New("block is nil")
			}
			fmt.Println("--------1")
			serializeBlock, err := Blockchain.SerializeBlock(v)
			fmt.Println("--------2")
			if err != nil {
				return err
			}
			fmt.Println("--------3")
			block := Blockchain.Chain{
				Id:    uuid.NewString(),
				Hash:  v.CurrHash,
				Block: serializeBlock,
			}
			arrayToMerge = append(arrayToMerge, &block)
			fmt.Println("--------4")
		}
	}
	Mutex.Lock()
	//var blocks []*Blockchain.Chain
	//dbCompare.Find(&blocks)
	errDelete := dbNode.Exec("DELETE FROM Chains WHERE Id != 0")
	if errDelete.Error != nil {
		return errDelete.Error
	}
	for _, v := range arrayToMerge {
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
	hash, err := Blockchain.LastHash("Start")
	if err != nil {
		log.Fatal(err)
	}
	BlockForTransaction, err = Blockchain.NewBlock(hash, "Start")
	if err != nil {
		log.Fatal(err)
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
