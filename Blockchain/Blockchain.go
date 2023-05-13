package Blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var (
	GenesisBlock string
	StorageChain string
)

const (
	TxsLimit = 4
)

//type BlockChain struct {
//	DB    *gorm.DB
//	index uint64
//}

func init() {
	viper.SetConfigFile("./LowConf/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	GenesisBlock = viper.GetString("GENESIS_BLOCK")
	StorageChain = viper.GetString("STORAGE_CHAIN")
}

func NewChain(VotesCount uint64, ChainMaster string) error {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	curTime, err := GetTime()
	var blocks []*Chain
	db.Find(&blocks)
	for _, v := range blocks {
		desBlock := DeserializeBlock(v.Block)
		if desBlock.ChainMaster == ChainMaster {
			return errors.New("affiliation already exist")
		}
	}
	genesis := &Block{
		PrevHash:    GenesisBlock,
		BalanceMap:  make(map[string]uint64),
		TimeStamp:   curTime,
		ChainMaster: ChainMaster,
	}
	genesis.BalanceMap[StorageChain] = VotesCount
	genesis.CurrHash = genesis.Hash()
	err = AddBlock(genesis)
	if err != nil {
		return err
	}
	return nil
}

func NewBlock(prevHash string, miner string, chainMaster string) (*Block, error) {
	VarConf := viper.GetString("DIFFICULTY")
	difficulty, err := strconv.Atoi(VarConf)
	curTime, err := GetTime()
	if err != nil {
		return nil, err
	}
	return &Block{
		Difficulty:  uint64(difficulty),
		PrevHash:    prevHash,
		BalanceMap:  make(map[string]uint64),
		Miner:       miner,
		ChainMaster: chainMaster,
		TimeStamp:   curTime,
	}, nil
}

func NewTransaction(
	fromUser *User,
	toUser string,
	lastHash string,
	value uint64) (*Transaction, error) {
	VarConf := viper.GetString("RAND_BYTES")
	VarConfConversion, err := strconv.Atoi(VarConf)
	randBytes, err := GenerateRandomBytes(uint64(VarConfConversion))
	if err != nil {
		return nil, err
	}
	tran := &Transaction{
		RandBytes: randBytes,
		PrevBlock: lastHash,
		Sender:    fromUser.Address(),
		Receiver:  toUser,
		Value:     value,
	}
	tran.CurrHash = tran.Hash()
	tran.Signature = tran.Sign(fromUser.Address() + toUser)
	return tran, err
}

func NewTransactionFromChain(
	master string,
	toUser string,
	lastHash string,
	value uint64) (*Transaction, error) {
	VarConf := viper.GetString("RAND_BYTES")
	VarConfConversion, err := strconv.Atoi(VarConf)
	randBytes, err := GenerateRandomBytes(uint64(VarConfConversion))
	if err != nil {
		return nil, err
	}
	tran := &Transaction{
		RandBytes: randBytes,
		PrevBlock: lastHash,
		Sender:    master,
		Receiver:  toUser,
		Value:     value,
	}
	tran.CurrHash = tran.Hash()
	tran.Signature = tran.Sign(master + toUser)
	return tran, err
}

func LastHash(master string) (string, error) {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return "", err
	}
	var hash string
	var chain []*Chain
	var blocks []*Block
	db.Find(&chain)
	for _, v := range chain {
		blocks = append(blocks, DeserializeBlock(v.Block))
	}
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].TimeStamp.AsTime().After(blocks[j].TimeStamp.AsTime())
	})
	for _, v := range blocks {
		fmt.Println(v)
		if v.ChainMaster == master {
			hash = v.CurrHash
			break
		}
	}
	return hash, nil
}

func Sign(privateKey string, data string) string {
	//tempSign := bytes.Join([][]byte{
	//	[]byte(privateKey),
	//	[]byte(data),
	//},
	//	[]byte{})
	tempSign := privateKey + data
	signature := HashSum(tempSign)
	return signature
}

func AddBlock(block *Block) error {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	errInsert := db.Exec("INSERT INTO Chains (Id, Hash, Block) VALUES ($1, $2, $3)",
		uuid.NewString(),
		block.CurrHash,
		SerializeBlock(block),
	)
	if errInsert.Error != nil {
		return errInsert.Error
	}
	return nil
}

// NewUser same with AddPass (BLCHxVote)
func NewUser(passport string) error {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	privateGenKey, err := GenerateKey()
	if err != nil {
		return err
	}
	db.Exec("INSERT INTO RelationPatterns (Id, PersonIdentifier, PrivateKeyTemplate) VALUES ($1, $2, $3)",
		uuid.NewString(),
		SetHash(passport),
		privateGenKey)
	return nil
}

// NewPublicKeyItem Same with NewUser(BLCHxVote)
func NewPublicKeyItem(affiliation string) (*User, error) {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	tempKey, err := GenerateKey()
	if err != nil {
		return nil, err
	}
	tempUUID := uuid.New()
	db.Exec("INSERT INTO PublicKeySets (Id, PublicKey, IsUsed, VotingAffiliation) VALUES ($1, $2, $3, $4)",
		tempUUID,
		tempKey,
		false,
		affiliation)
	return &User{
		Uuid:        tempUUID,
		PublicKey:   tempKey,
		IsUsed:      false,
		Affiliation: affiliation,
	}, nil
}

func NewCandidate(description string, affiliation string) (*Candidate, error) {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	tempUUID := uuid.New()
	tempKey, err := GenerateKey()
	db.Exec("INSERT INTO ElectionSubjects (Id, PublicKey,Description, VotingAffiliation) VALUES ($1, $2, $3, $4)",
		tempUUID,
		tempKey,
		description,
		affiliation)
	return &Candidate{
		Uuid:              tempUUID,
		PublicKey:         tempKey,
		Description:       description,
		VotingAffiliation: affiliation,
	}, nil
}

func Size(master string) (uint64, error) {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return 0, err
	}
	var index uint64
	var blocks []*Chain
	db.Find(&blocks)
	for _, v := range blocks {
		desBlock := DeserializeBlock(v.Block)
		if desBlock.ChainMaster == master {
			index++
		}
	}
	return index, nil
}

func Balance(address string, master string) (uint64, error) {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return 0, err
	}
	var balance uint64
	var blocks []*Chain
	//db.Where("Id = ?", size).Find(&blocks)
	db.Find(&blocks)
	for _, v := range blocks {
		desBlock := DeserializeBlock(v.Block)
		if desBlock.ChainMaster == master {
			if value, ok := desBlock.BalanceMap[address]; ok {
				balance = value
				break
			}
		}
	}
	return balance, nil
}

func (block *Block) Accept(user *User, master string, ch chan bool) error {
	curTime, err := GetTime()
	if err != nil {
		return err
	}
	block.TimeStamp = curTime
	block.CurrHash = block.Hash()
	privateKey, err := user.Private()
	if err != nil {
		return err
	}
	block.Signatures = block.Sign(privateKey)
	block.Nonce = block.Proof(ch)
	block.ChainMaster = master
	return nil
}

func GeneratePrivate(passport string, salt string, PublicKey string) (string, error) {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return "", err
	}
	var template string
	db.Raw("SELECT PrivateKeyTemplate FROM RelationPatterns WHERE PersonIdentifier = $1",
		SetHash(passport)).Scan(&template)
	if template == "" {
		return "", errors.New("identifier does not exist")
	}
	errExec := db.Exec("UPDATE PublicKeySets SET isUsed = 1 WHERE PublicKey = $1", PublicKey)
	if errExec.Error != nil {
		return "", errExec.Error
	}
	hash := sha256.Sum256([]byte(template + salt))
	err = ImportToDB(string(hash[:]), PublicKey)
	if err != nil {
		return "", err
	}
	return string(hash[:]), nil
}

func ImportToDB(PrivateKey string, PublicKey string) error {
	db, err := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	db.Exec("INSERT INTO KeyLinks (Id, PublicKey, PrivateKey) VALUES ($1, $2, $3)",
		uuid.NewString(),
		PublicKey,
		PrivateKey)
	return nil
}

// GenerateKey TODO Rewrite
func GenerateKey() (string, error) {
	TimeUrl := viper.GetString("TIME_URL")
	resp, err := http.Get(TimeUrl)
	if err != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(v.GetStringBytes("dateTime"))
	return string(hash[:]), nil
}
