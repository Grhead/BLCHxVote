package Blockchain

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fastjson"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"time"
)

type BlockChain struct {
	DB    *sql.DB
	index uint64
}

type Block struct {
	CurrHash     []byte
	PrevHash     []byte
	TimeStamp    string
	Transactions []Transaction
	Mapping      map[string]uint64
}

type Transaction struct {
	RandBytes []byte
	PrevBlock []byte
	Sender    string
	Receiver  string
	Value     uint64
	Signature []byte
	CurrHash  []byte
}

type User struct {
	PublicKey string
}
type Candidate struct {
	PublicKey   string
	Description string
}

const (
	CREATE_TABLE = `CREATE TABLE BlockChain (
	 					Id INTEGER PRIMARY KEY AUTOINCREMENT,
						Hash VARCHAR(44) UNIQUE,
	 					Block TEXT
						);`
	CREATE_PASSDB = `CREATE TABLE TemplateDB (
	 					Id INTEGER PRIMARY KEY AUTOINCREMENT,
						Passport TEXT,
	 					TemplatePRK TEXT
						);`
	CREATE_PAREDB = `CREATE TABLE Pare (
	 					Id INTEGER PRIMARY KEY AUTOINCREMENT,
						PrivateK TEXT UNIQUE,
	 					PublicK TEXT UNIQUE
						);`
	CREATE_PUBLICDB = `CREATE TABLE PublicDB (
	 					Id INTEGER PRIMARY KEY AUTOINCREMENT,
	 					PublicK TEXT UNIQUE
						);`
)

const (
	GENESIS_BLOCK = "Grenka"
	STORAGE_VALUE = 100
	STORAGE_CHAIN = "GRChain"
	RAND_BYTES    = 32
	TXS_LIMIT     = 2
	TIME_URL      = "http://worldtimeapi.org/api/ip"
)

var (
	resp, _ = http.Get(TIME_URL)
)

func NewVotePass(Pasefile string, Parefile string, Publicfile string) error {
	file, err := os.Create(Pasefile)
	if err != nil {
		return err
	}
	file.Close()

	dbpass, err := sql.Open("sqlite3", Pasefile)
	if err != nil {
		return err
	}
	defer dbpass.Close()
	_, err = dbpass.Exec(CREATE_PASSDB)

	file, err = os.Create(Parefile)
	if err != nil {
		return err
	}
	file.Close()

	dbpare, err := sql.Open("sqlite3", Parefile)
	if err != nil {
		return err
	}
	defer dbpare.Close()
	_, err = dbpare.Exec(CREATE_PAREDB)

	file, err = os.Create(Publicfile)
	if err != nil {
		return err
	}
	file.Close()

	dbpublic, err := sql.Open("sqlite3", Publicfile)
	if err != nil {
		return err
	}
	defer dbpublic.Close()
	_, err = dbpublic.Exec(CREATE_PUBLICDB)

	return nil
}

func NewChain(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	file.Close()

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}

	defer db.Close()
	_, err = db.Exec(CREATE_TABLE)
	if err != nil {
		return nil
	}
	chain := &BlockChain{
		DB: db,
	}
	genesis := &Block{
		PrevHash:  []byte(GENESIS_BLOCK),
		Mapping:   make(map[string]uint64),
		TimeStamp: time.Now().Format(time.RFC3339),
	}
	genesis.Mapping[STORAGE_CHAIN] = STORAGE_VALUE
	genesis.CurrHash = genesis.Hash()
	chain.AddBlock(genesis)
	return nil
}
func GetTokens(receiver *User, chain *BlockChain, value uint64) {
	block := NewBlock(chain.LastHash())
	block.AddTransaction(chain, &Transaction{
		RandBytes: GenerateRandomBytes(RAND_BYTES),
		PrevBlock: chain.LastHash(),
		Sender:    STORAGE_CHAIN,
		Receiver:  receiver.Address(),
		Value:     value,
	})
	block.Accept(chain)
	chain.AddBlock(block)
}
func LoadChain(filename string) *BlockChain {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil
	}
	chain := &BlockChain{
		DB: db,
	}
	chain.index = chain.Size()
	return chain
}

func (chain *BlockChain) Size() uint64 {
	var index uint64
	row := chain.DB.QueryRow("SELECT Id FROM BlockChain ORDER BY Id DESC")
	row.Scan(&index)
	return index
}

func (chain *BlockChain) AddBlock(block *Block) {
	chain.index += 1
	chain.DB.Exec("INSERT INTO BlockChain (Hash, Block) VALUES ($1, $2)",
		Base64Encode(block.CurrHash),
		SerializeBlock(block),
	)
}

func NewBlock(prevHash []byte) *Block {
	return &Block{
		PrevHash: prevHash,
		Mapping:  make(map[string]uint64),
	}
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func SerializeBlock(block *Block) string {
	jsonData, err := json.MarshalIndent(*block, "", "\t")
	if err != nil {
		return ""
	}
	return string(jsonData)
}
func DeserializeBlock(data string) *Block {
	var block Block
	err := json.Unmarshal([]byte(data), &block)
	if err != nil {
		return nil
	}
	return &block
}

func (tran *Transaction) Hash() []byte {
	return HashSum(bytes.Join(
		[][]byte{
			tran.RandBytes,
			tran.PrevBlock,
			[]byte(tran.Sender),
			[]byte(tran.Receiver),
			ToBytes(tran.Value),
		},
		[]byte{},
	))
}

func (block *Block) Hash() []byte {
	var tempHash []byte
	for _, tx := range block.Transactions {
		tempHash = HashSum(bytes.Join(
			[][]byte{
				tempHash,
				tx.CurrHash,
			},
			[]byte{},
		))
	}
	var list []string
	for hash := range block.Mapping {
		list = append(list, hash)
	}
	sort.Strings(list)
	for _, hash := range list {
		tempHash = HashSum(bytes.Join(
			[][]byte{
				tempHash,
				[]byte(hash),
				ToBytes(block.Mapping[hash]),
			},
			[]byte{},
		))
	}

	return HashSum(bytes.Join(
		[][]byte{
			tempHash,
			block.PrevHash,
			[]byte(block.TimeStamp),
		},
		[]byte{},
	))
}

func ToBytes(data uint64) []byte {
	var buf = new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func HashSum(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

//func NewTransaction(user *User, toUser string, lastHash []byte, value uint64) *Transaction {
//	tran := &Transaction{
//		RandBytes: GenerateRandomBytes(RAND_BYTES),
//		PrevBlock: lastHash,
//		Sender:    user.Address(),
//		Receiver:  toUser,
//		Value:     value,
//	}
//	tran.CurrHash = tran.Hash()
//	tran.Signature = tran.Sign(user.Private())
//	return tran
//}

func GenerateRandomBytes(max uint64) []byte {
	var slice = make([]byte, max)
	_, err := rand.Read(slice)
	if err != nil {
		return nil
	}
	return slice
}

func (tran *Transaction) Sign(privk []byte) []byte {
	return Sign(privk, tran.CurrHash)
}

func Sign(privk []byte, data []byte) []byte {
	tempmeta := bytes.Join([][]byte{privk, data}, []byte{})
	signature := HashSum(tempmeta)
	return signature
}

func (block *Block) AddTransaction(chain *BlockChain, tran *Transaction) error {
	var balanceInChain uint64
	balanceInTX := tran.Value
	if value, ok := block.Mapping[tran.Sender]; ok {
		balanceInChain = value
	} else {
		balanceInChain = chain.Balance(tran.Sender, chain.Size())
	}
	if balanceInTX > balanceInChain {
		return errors.New("insufficient funds")
	}
	block.Mapping[tran.Sender] = balanceInChain - balanceInTX
	block.addBalance(chain, tran.Receiver, tran.Value)
	block.Transactions = append(block.Transactions, *tran)
	return nil
}

func (block *Block) addBalance(chain *BlockChain, receiver string, value uint64) {
	var balanceInChain uint64
	if v, ok := block.Mapping[receiver]; ok {
		balanceInChain = v
	} else {
		balanceInChain = chain.Balance(receiver, chain.Size())
	}
	block.Mapping[receiver] = balanceInChain + value
}

func (chain *BlockChain) Balance(address string, size uint64) uint64 {
	var balance uint64
	var sblock string
	var block *Block
	rows, err := chain.DB.Query("SELECT Block FROM BlockChain WHERE Id <= $1 ORDER BY ID DESC", size)
	if err != nil {
		return balance
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&sblock)
		block = DeserializeBlock(sblock)
		if value, ok := block.Mapping[address]; ok {
			balance = value
			break
		}
	}
	return balance
}

func (block *Block) Accept(chain *BlockChain) error {
	block.TimeStamp = time.Now().Format(time.RFC3339)
	block.CurrHash = block.Hash()
	return nil
}

func DeserializeTX(data string) *Transaction {
	var tx Transaction
	err := json.Unmarshal([]byte(data), &tx)
	if err != nil {
		return nil
	}
	return &tx
}
func (chain *BlockChain) LastHash() []byte {
	var hash string
	row := chain.DB.QueryRow("SELECT Hash FROM BlockChain ORDER BY Id DESC")
	row.Scan(&hash)
	return Base64Decode(hash)
}

func Base64Decode(data string) []byte {
	result, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil
	}
	return result
}

func NewUser(filename string) *User {
	var temp string = GenerateKey()
	db, _ := sql.Open("sqlite3", filename)
	db.Exec("INSERT INTO PublicDB (PublicK) VALUES ($1)", temp)
	return &User{
		PublicKey: temp,
	}
}

func LoadUser(privateK string, filename string) *User {
	db, _ := sql.Open("sqlite3", filename)
	var result string
	db.QueryRow("SELECT PublicK FROM Pare WHERE PrivateK = $1", privateK).Scan(&result)
	defer db.Close()
	if result == "" {
		return nil
	}
	return &User{
		PublicKey: result,
	}
}

func Purse(passport string, filename string) error {
	db, _ := sql.Open("sqlite3", filename)
	var result string
	db.QueryRow("SELECT TemplatePRK FROM TemplateDB WHERE Passport = $1", passport).Scan(&result)
	defer db.Close()
	if result == "" {
		return errors.New("Empty select")
	}
	return nil
}

func AddPass(passport string, filename string) error {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}
	db.Exec("INSERT INTO TemplateDB (Passport, TemplatePRK) VALUES ($1, $2)",
		passport,
		GenerateKey(),
	)
	defer db.Close()
	return nil
}
func Private(passport string, salt string, TemplateDBFile string, PareDBFile string, PKey string) string {
	templateDB, _ := sql.Open("sqlite3", TemplateDBFile)
	var template string
	templateDB.QueryRow("SELECT TemplatePRK FROM TemplateDB WHERE Passport = $1", passport).Scan(&template)
	defer templateDB.Close()
	if template == "" {
		//return errors.New("empty pass cell")
		return "Empty"
	}
	hash := sha256.New()
	hash.Write([]byte(template + salt))
	result := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("resylt %1", result)
	ImportToDB(PareDBFile, result, PKey)
	return result
}
func ImportToDB(PareDBFile string, PrKResult string, PKey string) error {
	PareDB, _ := sql.Open("sqlite3", PareDBFile)
	PareDB.Exec("INSERT INTO Pare (PrivateK, PublicK) VALUES ($1, $2)", PrKResult, PKey)
	defer PareDB.Close()
	return nil
}
func (user *User) Address() string {
	return user.PublicKey
}

func GenerateKey() string {
	//defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var p fastjson.Parser
	v, _ := p.Parse(string(body))
	hash := sha256.New()
	hash.Write(v.GetStringBytes("datetime"))
	priv := hex.EncodeToString(hash.Sum(nil))
	//t, _ := hex.DecodeString(priv)
	return priv
}
