package main

import (
	"bytes"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	bc "BLCHxVote/Blockchain"
	nt "BLCHxVote/Network"
)

var (
	Filename string
	Serve    string
	Chain    *bc.BlockChain
	Block    *bc.Block
)

func init() {
	if len(os.Args) < 2 {
		panic("failed: len(os.Args) < 2")
	}
	var (
		serveStr       = ""
		addrStr        = ""
		chainNewStr    = ""
		chainLoadStr   = ""
		chainNewNumStr = ""
		nums           uint64
	)
	var (
		chainNewExist  = false
		chainLoadExist = false
	)
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch {
		case strings.HasPrefix(arg, "-serve:"):
			serveStr = strings.Replace(arg, "-serve:", "", 1)
			//serveExist = true
		case strings.HasPrefix(arg, "-loadaddr:"):
			addrStr = strings.Replace(arg, "-loadaddr:", "", 1)
			//addrExist = true
		case strings.HasPrefix(arg, "-newchain:"):
			chainNewStr = strings.Replace(arg, "-newchain:", "", 1)
			chainNewExist = true
		case strings.HasPrefix(arg, "-count:"):
			chainNewNumStr = strings.Replace(arg, "-count:", "", 1)
			temp, _ := strconv.Atoi(chainNewNumStr)
			nums = uint64(temp)
			chainNewExist = true
		case strings.HasPrefix(arg, "-loadchain:"):
			chainLoadStr = strings.Replace(arg, "-loadchain:", "", 1)
			chainLoadExist = true
		}
	}
	Serve = serveStr
	var addresses []string
	err := json.Unmarshal([]byte(readFile(addrStr)), &addresses)
	if err != nil {
		panic("failed: load addresses")
	}
	var mapaddr = make(map[string]bool)
	for _, addr := range addresses {
		if addr == Serve {
			continue
		}
		if _, ok := mapaddr[addr]; ok {
			continue
		}
		mapaddr[addr] = true
		Addresses = append(Addresses, addr)
	}
	if chainNewExist {
		Filename = chainNewStr
		Chain = chainNew(chainNewStr, nums)
	}
	if chainLoadExist {
		Filename = chainLoadStr
		Chain = chainLoad(chainLoadStr)
	}
	if Chain == nil {
		panic("failed: load chain")
	}
	Block = bc.NewBlock(Chain.LastHash())
}

func main() {
	nt.Listen(Serve, handleServer)
	for {
		fmt.Scanln()
	}
}

func chainNew(filename string, num uint64) *bc.BlockChain {
	err := bc.NewChain(filename, num)
	if err != nil {
		return nil
	}
	return bc.LoadChain(filename)
}
func chainLoad(filename string) *bc.BlockChain {
	chain := bc.LoadChain(filename)
	return chain
}
func userNew(filename string) *bc.User {
	user := bc.NewUser(filename)
	if user == nil {
		return nil
	}
	return user
}
func userLoad(privateK string) *bc.User {
	user := bc.LoadUser(privateK, "Databases/paredb.db")
	if user == nil {
		return nil
	}
	return user
}

func handleServer(conn nt.Conn, pack *nt.Package) {
	nt.Handle(ADD_BLOCK, conn, pack, addBlock)
	nt.Handle(ADD_TRNSX, conn, pack, addTransaction)
	nt.Handle(GET_BLOCK, conn, pack, getBlock)
	nt.Handle(GET_LHASH, conn, pack, getLastHash)
	nt.Handle(GET_BLNCE, conn, pack, getBalance)
	nt.Handle(GET_CSIZE, conn, pack, getChainSize)
}

func compareChains(address string, num uint64) {
	fmt.Println("Compare")
	filename := "temp_" + hex.EncodeToString(bc.GenerateRandomBytes(32))
	file, _ := os.Create(filename)
	file.Close()
	defer func() {
		os.Remove(filename)
	}()
	res := nt.Send(address, &nt.Package{
		Option: GET_BLOCK,
		Data:   fmt.Sprintf("%d", 0),
	})
	genesis := bc.DeserializeBlock(res.Data)
	if !bytes.Equal(genesis.CurrHash, hashBlock(genesis)) {
		return
	}
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec(bc.CREATE_TABLE)
	chain := &bc.BlockChain{
		DB: db,
	}
	chain.AddBlock(genesis)
	for i := uint64(1); i < num; i++ {
		res := nt.Send(address, &nt.Package{
			Option: GET_BLOCK,
			Data:   fmt.Sprintf("%d", i),
		})
		block := bc.DeserializeBlock(res.Data)
		chain.AddBlock(block)
	}
	//Mutex.Lock()
	Chain.DB.Close()
	os.Remove(Filename)
	copyFile(filename, Filename)
	Chain = bc.LoadChain(Filename)
	Block = bc.NewBlock(Chain.LastHash())
	//Mutex.Unlock()
	return
}

func pushBlockToNet(block *bc.Block) {
	var sblock = bc.SerializeBlock(block)
	var msg = Serve + SEPARATOR + fmt.Sprintf("%d", Chain.Size()) + SEPARATOR + sblock
	for _, addr := range Addresses {
		go nt.Send(addr, &nt.Package{
			Option: ADD_BLOCK,
			Data:   msg,
		})
	}
}

func addBlock(pack *nt.Package) string {
	splited := strings.Split(pack.Data, SEPARATOR)
	block := bc.DeserializeBlock(splited[2])
	currSize := Chain.Size()
	num, _ := strconv.Atoi(splited[1])
	if currSize < uint64(num) {
		compareChains(splited[0], uint64(num))
		return "ok "
	}
	//Mutex.Lock()
	Chain.AddBlock(block)
	Block = bc.NewBlock(Chain.LastHash())
	//Mutex.Unlock()
	return "ok"
}

func addTransaction(pack *nt.Package) string {
	var tx = bc.DeserializeTX(pack.Data)
	//Mutex.Lock()
	Block.AddTransaction(Chain, tx)
	//Mutex.Unlock()

	go func() {
		//Mutex.Lock()
		block := *Block
		//Mutex.Unlock()
		res := (&block).Accept()
		//Mutex.Lock()
		if res == nil && bytes.Equal(block.PrevHash, Block.PrevHash) {
			Chain.AddBlock(&block)
			pushBlockToNet(&block)
		}
		Block = bc.NewBlock(Chain.LastHash())
		//Mutex.Unlock()
	}()
	return "ok"
}

func getBlock(pack *nt.Package) string {
	num, err := strconv.Atoi(pack.Data)
	if err != nil {
		return ""
	}
	size := Chain.Size()
	if uint64(num) < size {
		return selectBlock(Chain, num)
	}
	return ""
}
func getLastHash(pack *nt.Package) string {
	fmt.Println("getLH")
	return bc.Base64Encode(Chain.LastHash())
}
func getBalance(pack *nt.Package) string {
	return fmt.Sprintf("%d", Chain.Balance(pack.Data, Chain.Size()))
}
func getChainSize(pack *nt.Package) string {
	return fmt.Sprintf("%d", Chain.Size())
}
func selectBlock(chain *bc.BlockChain, i int) string {
	fmt.Println("SelectBL")
	var block string
	row := chain.DB.QueryRow("SELECT Block FROM BlockChain WHERE Id=$1", i+1)
	row.Scan(&block)
	return block
}
func hashBlock(block *bc.Block) []byte {
	var tempHash []byte
	for _, tx := range block.Transactions {
		tempHash = bc.HashSum(bytes.Join(
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
		tempHash = bc.HashSum(bytes.Join(
			[][]byte{
				tempHash,
				[]byte(hash),
				bc.ToBytes(block.Mapping[hash]),
			},
			[]byte{},
		))
	}
	return bc.HashSum(bytes.Join(
		[][]byte{
			tempHash,
			block.PrevHash,
			[]byte(block.TimeStamp),
		},
		[]byte{},
	))
}
func copyFile(src, dst string) error {
	fmt.Println("copy")
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
