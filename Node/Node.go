package main

import (
	"VOX2/Blockchain"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fastjson"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type ChainHelp struct {
	Master string `form:"master" json:"master"`
	Count  uint64 `form:"count" json:"count"`
}
type MasterHelp struct {
	Master string `form:"master" json:"master"`
}
type UserHelp struct {
	User string `form:"user" json:"user"`
}
type BlockHelp struct {
	Block   *Blockchain.Block `form:"user" json:"block"`
	Size    uint64            `form:"size" json:"size"`
	Address string            `form:"address" json:"address"`
}

var Mutex sync.Mutex
var IsMining bool
var BreakMining = make(chan bool)
var Block *Blockchain.Block
var ThisServe string
var OtherAddresses []*fastjson.Value

// TODO --CreateNewChain	// TODO NewChain
// TODO --CompareChains		// TODO NewBlock
// TODO --PushBlockToNet	// TODO NewTransaction
// TODO --AddBlock			// TODO NewTransactionFromChain
// TODO --AddTransaction	// TODO --LastHash
// TODO --GetBlock_const	// TODO AddBlock
// TODO --GetLastHash		// TODO NewDormantUser
// TODO --GetBalance		// TODO LoadToEnterAlreadyUser
// TODO --GetChainSize		// TODO NewPublicKeyItem
// TODO __SelectBlock		// TODO NewCandidate
// TODO __HashBlock			// TODO Size
// TODO __CopyFile			// TODO Balance

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
}

func main() {
	router := gin.Default()

	router.POST("/newchain", GinNewChain)
	router.POST("/addblock", GinAddBlock)
	router.POST("/addtx", GinAddTransaction)
	router.POST("/getblock", GinGetBlocks)
	router.POST("/getlasthash", GinGetLastHash)
	router.POST("/getbalance", GinGetBalance)
	router.POST("/getchainsize", GinGetChainSize)

	router.POST("/netpush", GinPushBlockToNet)

	err := router.Run(strings.Trim(ThisServe, "\""))
	if err != nil {
		panic(err)
		return
	}
}

func GinNewChain(c *gin.Context) {
	var input *ChainHelp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		genesisHash, errChain := NewChain(input.Master, input.Count)
		if errChain != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"GenesisHash": genesisHash})
	}
}

func GinAddBlock(c *gin.Context) {
	var input *BlockHelp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		block, errAdd := AddBlock(input)
		if errAdd != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"BlockAddStatus": block})
	}
}

func GinAddTransaction(c *gin.Context) {
	var input *Blockchain.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		transaction, err := AddTransaction(input)
		if err != nil {
			return
		}
		c.JSON(200, gin.H{"AddTxStatus": transaction})
	}
}
func GinGetBlocks(c *gin.Context) {
	var input *MasterHelp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		blocks, errGet := GetBlocks(input)
		if errGet != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blocks)
	}
}
func GinGetLastHash(c *gin.Context) {
	var input *MasterHelp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		hash, errLH := GetLastHash(input)
		if errLH != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"LastHash": hash})
	}
}
func GinGetBalance(c *gin.Context) {
	var input *UserHelp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		balance, errBalance := GetBalance(input)
		if errBalance != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"Balance": balance})
	}
}
func GinGetChainSize(c *gin.Context) {
	var input *MasterHelp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		size, errChainSize := GetChainSize(input)
		if errChainSize != nil {
			return
		}
		c.JSON(200, gin.H{"ChainSize": size})
	}
}
func GinPushBlockToNet(c *gin.Context) {
	var input *Blockchain.Block
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		errBlock := PushBlockToNet(input)
		if errBlock != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"Status": "ok"})
	}
}
