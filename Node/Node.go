package main

import (
	"VOX2/Blockchain"
	"errors"
	"github.com/gin-contrib/cors"
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
	Block   *Blockchain.Block `form:"block" json:"block"`
	Size    uint64            `form:"size" json:"size"`
	Address string            `form:"address" json:"address"`
}
type TransactionHelp struct {
	Master string                  `form:"master" json:"master"`
	Tx     *Blockchain.Transaction `form:"transaction" json:"transaction"`
}

var Mutex sync.Mutex
var IsMining bool
var BreakMining = make(chan bool)
var ThisServe string
var OtherAddresses []*fastjson.Value
var BlockForTransaction *Blockchain.Block

func SetAddress(server string) {
	ThisServe = ":" + server
}

func init() {
	hash, err := Blockchain.LastHash("Start")
	if err != nil {
		log.Fatalln(err)
	}
	BlockForTransaction, err = Blockchain.NewBlock(hash, "Start")
	if err != nil {
		log.Fatalln(err)
	}
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
	for i, v := range OtherAddresses {
		if strings.Contains(v.String(), ThisServe) {
			func(slice []*fastjson.Value, s int) []*fastjson.Value {
				return append(slice[:s], slice[s+1:]...)
			}(OtherAddresses, i)
		}
	}
}

func main() {
	rootCmd.Flags().StringVarP(&option, "set", "s", "", "Set address")
	Execute()
	if ThisServe == "" {
		panic(errors.New("empty address"))
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,
	}))
	router.POST("/newchain", GinNewChain)
	router.POST("/addblock", GinAddBlock)
	router.POST("/addtx", GinAddTransaction)
	router.POST("/getblock", GinGetBlocks)
	router.POST("/getlasthash", GinGetLastHash)
	router.POST("/getbalance", GinGetBalance)
	router.POST("/getchainsize", GinGetChainSize)
	router.GET("/getdb", GinGetDb)

	router.POST("/netpush", GinPushBlockToNet)
	err := router.Run(strings.Trim(ThisServe, "\""))
	if err != nil {
		log.Fatalln(err)
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
	var input *TransactionHelp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		transaction, errTx := AddTransaction(input)
		if errTx != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": errTx.Error()})
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
func GinGetDb(c *gin.Context) {
	db, err := Blockchain.GetFullDb()
	if err != nil {
		return
	}
	c.JSON(200, db)
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
	var input *BlockHelp
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	} else {
		errBlock := pushBlockToNet(input)
		if errBlock != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"Status": "ok"})
	}
}
