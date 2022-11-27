package Blockchain

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
	PublicKey []byte
}

const (
	CREATE_TABLE = `CREATE TABLE BlockChain (
	 					Id INTEGER PRIMARY KEY AUTOINCREMENT,
						Hash VARCHAR(44) UNIQUE,
	 					Block TEXT
						);`
)

const (
	GENESIS_BLOCK = "Grenka"
	STORAGE_VALUE = 100
	STORAGE_CHAIN = "GRChain"
	RAND_BYTES    = 32
	TXS_LIMIT     = 2
)

func NewChain(filename string, receiver string) error {
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
	genesis.Mapping[receiver] = 50
	chain.AddBlock(genesis)
	return nil
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

func NewTransaction(user *User, toUser string, lastHash []byte, value uint64) *Transaction {
	tran := &Transaction{
		RandBytes: GenerateRandomBytes(RAND_BYTES),
		PrevBlock: lastHash,
		Sender:    user.Address(),
		Receiver:  toUser,
		Value:     value,
	}
	tran.CurrHash = tran.Hash()
	tran.Signature = tran.Sign(user.Private([]byte("hello")))
	return tran
}

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

func (user *User) Address() string {
	return string(user.PublicKey[:])
}

func (user *User) Private(salt []byte) []byte {
	tempPrivate := (bytes.Join(
		[][]byte{
			user.PublicKey,
			salt,
		},
		[]byte{},
	))
	return tempPrivate
}

func (block *Block) AddTransaction(chain *BlockChain, tran *Transaction) error {
	// if tran == nil {
	// 	return errors.New("tran is null")
	// }
	// if tran.Value == 0 {
	// 	return errors.New("tran value = 0")
	// }
	// if tran.Sender != STORAGE_CHAIN && len(block.Transactions) == TXS_LIMIT {
	// 	return errors.New("len tran = limit")
	// }
	// if !bytes.Equal(tran.PrevBlock, chain.LastHash()) {
	// 	return errors.New("prev block in tran /= last hash in chain")
	// }
	var balanceInChain uint64
	balanceInTX := tran.Value
	if value, ok := block.Mapping[tran.Sender]; ok {
		balanceInChain = value
	} else {
		balanceInChain = chain.Balance(tran.Sender, chain.Size())
	}
	// if balanceInTX > balanceInChain {
	// 	return errors.New("not enought")
	// }
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
	rows, err := chain.DB.Query("SELECT Block FROM BlockChain WHERE Id <= $1 ORDER BY ID DESC",
		chain.index)
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

func (block *Block) Accept(chain *BlockChain, user *User, ch chan bool) error {
	// if !block.transactionsIsValid(chain, chain.Size()) {
	// 	return errors.New("tran is not valid")
	// }
	block.TimeStamp = time.Now().Format(time.RFC3339)
	block.CurrHash = block.Hash()
	return nil
}

func SerializeTX(tx *Transaction) string {
	jsonData, err := json.MarshalIndent(*tx, "", "\t")
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func DeserializeTX(data string) *Transaction {
	var tx Transaction
	err := json.Unmarshal([]byte(data), &tx)
	if err != nil {
		return nil
	}
	return &tx
}

func NewUser() *User {
	return &User{
		PublicKey: GeneratePublic(),
	}
}
func GeneratePublic() []byte {
	// priv, err := rsa.GenerateKey(rand.Reader, int(bits))
	rand.Seed(time.Now().UnixNano())
	priv := bytes.Join(
		[][]byte{
			[]byte(strconv.FormatInt(int64(rand.Int63()),
				rand.Intn(10))),
		},
		[]byte{},
	)
	return priv
}

func LoadUser(purse string) *User {
	priv := []byte(purse)
	if priv == nil {
		return nil
	}
	return &User{
		PublicKey: priv,
	}
}

func (user *User) Purse() string {
	return string((user.Private([]byte("hello"))))
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
