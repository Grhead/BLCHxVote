package Basic

import (
	"VOX2/Blockchain"
	"VOX2/Transport"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-co-op/gocron"
	_ "github.com/go-co-op/gocron"
	"github.com/imroc/req/v3"
	"github.com/valyala/fastjson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var MiningResponse Transport.CheckHelp

func CallCreateVoters(voter interface{}, master string) ([]*Blockchain.User, error) {
	var resultItems []*Blockchain.User
	switch voter.(type) {
	case int:
		max, err := strconv.Atoi(fmt.Sprintf("%v", voter))
		if err != nil {
			return nil, err
		}
		for i := 0; i < max; i++ {
			item, errNewPublicKey := Blockchain.NewPublicKeyItem(master)
			if errNewPublicKey != nil {
				return nil, errNewPublicKey
			}
			resultItems = append(resultItems, item)
		}
	case string:
		err := Blockchain.NewDormantUser(fmt.Sprintf("%v", voter))
		if err != nil {
			return nil, err
		}
		item, err := Blockchain.NewPublicKeyItem(master)
		if err != nil {
			return nil, err
		}
		resultItems = append(resultItems, item)
	default:
		return nil, errors.New("invalid type")
	}
	return resultItems, nil
}

func CallViewCandidates() ([]*Blockchain.ElectionSubjects, error) {
	db, err := gorm.Open(sqlite.Open("Database/ContractDB.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var Candidates []*Blockchain.ElectionSubjects
	db.Table("ElectionSubjects").Find(&Candidates)
	return Candidates, nil
}

func CallNewCandidate(description string, affiliation string) (*Blockchain.ElectionSubjects, error) {
	candidate, err := Blockchain.NewCandidate(description, affiliation)
	if err != nil {
		return nil, err
	}
	return candidate, nil
}

func GetBalance(userAddress string) (*Transport.BalanceHelp, error) {
	addresses, err := ReadAddresses()
	if err != nil {
		return nil, err
	}
	userAddressStruct := Transport.UserHelp{
		User: userAddress,
	}
	var userBalance Transport.BalanceHelp
	client := req.C().DevMode()
	for _, addr := range addresses {
		resp, errReq := client.R().
			SetBody(&userAddressStruct).
			SetSuccessResult(&userBalance).
			Post(fmt.Sprintf("http://%s/getbalance", strings.Trim(addr.String(), "\"")))
		if errReq != nil && !strings.Contains(errReq.Error(), "No connection could be made because the target machine actively refused it.") {
			return nil, errReq
		}
		if errReq == nil {
			if resp.Body == nil {
				continue
			}
		}
	}
	return &userBalance, nil
}

func ChainSize(master string) (string, error) {
	addresses, err := ReadAddresses()
	if err != nil {
		return "", err
	}
	masterChain := Transport.MasterHelp{Master: master}
	var chainSize Transport.SizeHelp
	client := req.C().DevMode()
	resp, err := client.R().
		SetBody(&masterChain).
		SetSuccessResult(&chainSize).
		Post(fmt.Sprintf("http://%s/getchainsize", strings.Trim(addresses[0].String(), "\"")))
	if err != nil {
		return "", err
	}
	if resp.Body == nil {
		return "", errors.New("empty response")
	}
	return chainSize.ChainSize, nil
}

func GetPartOfChain(master string) ([]*Blockchain.Block, error) {
	addresses, err := ReadAddresses()
	if err != nil {
		return nil, err
	}
	var partOfChain []*Blockchain.Block
	ChainMaster := Transport.MasterHelp{
		Master: master,
	}
	client := req.C().DevMode()
	resp, err := client.R().
		SetBody(&ChainMaster).
		SetSuccessResult(&partOfChain).
		Post(fmt.Sprintf("http://%s/getblock", strings.Trim(addresses[0].String(), "\"")))
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, errors.New("empty response")
	}
	return partOfChain, nil
}

func GetFullChain() ([]*Blockchain.Block, error) {
	addresses, err := ReadAddresses()
	if err != nil {
		return nil, err
	}
	var fullChain []*Blockchain.Block
	client := req.C().DevMode()
	resp, err := client.R().SetSuccessResult(&fullChain).
		Get(fmt.Sprintf("http://%s/getdb", strings.Trim(addresses[0].String(), "\"")))
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, errors.New("empty response")
	}
	return fullChain, nil
}

// AcceptNewUser TODO add time verification
func AcceptNewUser(Pass string, salt string, PublicKey string) (string, error) {
	//TODO add time verification
	//t, _ := time.ParseDuration(EndTime)
	//t1, _ := time.ParseDuration(LimitTime())
	//if t1 > t {
	//	return "time"
	//}
	private, err := Blockchain.RegisterGeneratePrivate(Pass, salt, PublicKey)
	if err != nil {
		return "", err
	}
	return private, nil
}

// AcceptLoadUser TODO add time verification
func AcceptLoadUser(PublicK string, PrivateK string) (*Blockchain.User, error) {
	//TODO add time verification
	//t, _ := time.ParseDuration(EndTime)
	//t1, _ := time.ParseDuration(LimitTime())
	//if t1 > t {
	//	return "2"
	//}
	UserPrivate, err := Blockchain.LoadToEnterAlreadyUserPrivate(PrivateK)
	if err != nil {
		return nil, err
	}
	UserPublic, err := Blockchain.LoadToEnterAlreadyUserPublic(PublicK)
	if err != nil {
		return nil, err
	}
	if UserPublic.PublicKey != PublicK || UserPrivate.PublicKey != PublicK || !reflect.DeepEqual(UserPrivate, UserPublic) {
		return nil, errors.New("invalid input")
	}
	bal, err := GetBalance(UserPublic.Address())
	if err != nil {
		return nil, err
	}
	if bal == nil {
		return nil, errors.New("zero balance")
	}
	return UserPublic, nil
}

func ChainTXBlock(receiver string, master string, num uint64) (string, error) {
	db, errDb := gorm.Open(sqlite.Open("Database/ContractDB.db"), &gorm.Config{})
	if errDb != nil {
		return "", errDb
	}
	addresses, err := ReadAddresses()
	if err != nil {
		return "", err
	}
	client := req.C().DevMode()
	var lastHash Transport.LastHashHelp
	var txStatus Transport.TransactionResponseHelp
	var transactionToNet Transport.TransactionHelp
	RequestData := Transport.MasterHelp{
		Master: master,
	}
	type TransactionsDb struct {
		Id           int
		Transactions string
	}

	var TransactionsArray []TransactionsDb
	var errNode error

	s1 := gocron.NewScheduler(time.UTC)
	_, errTask := s1.Every(10).Seconds().Do(SchedulerTask, client, addresses[1])
	if errTask != nil {
		return "", errTask
	}
	s1.StartAsync()

	for _, addr := range addresses {
		//_, errNode = client.R().SetSuccessResult(&MiningResponse).
		//	Get(fmt.Sprintf("http://%s/check", strings.Trim(addr.String(), "\"")))
		//if errNode != nil {
		//	if strings.Contains(errNode.Error(), "No connection could be made because the target machine actively refused it.") {
		//		fmt.Println("aaaaaaaaaaaaa")
		//		continue
		//	}
		//}
		resp, errReq := client.R().
			SetBody(&RequestData).
			SetSuccessResult(&lastHash).
			Post(fmt.Sprintf("http://%s/getlasthash", strings.Trim(addr.String(), "\"")))
		if errReq != nil && !strings.Contains(errReq.Error(), "No connection could be made because the target machine actively refused it.") {
			return "", errReq
		}
		if errReq == nil {
			if resp.Body == nil {
				continue
			}
		}
		balance, errBalance := GetBalance(master)
		if errBalance != nil {
			return "", errBalance
		}
		chainBalance, errConversion := strconv.Atoi(balance.Balance)
		if errConversion != nil {
			return "", errConversion
		}
		if uint64(chainBalance) < num {
			return "", errors.New("not enough chain founds")
		}
		public, errLoad := Blockchain.LoadToEnterAlreadyUserPublic(receiver)
		if errLoad != nil {
			return "", errLoad
		}
		tx, errNewTx := Blockchain.NewTransactionFromChain(master, public, num)
		if errNewTx != nil {
			return "", errNewTx
		}
		transactionToNet = Transport.TransactionHelp{
			Master: master,
			Tx:     tx,
		}
	}
	db.Raw("SELECT Id, Transactions FROM TransactionQueue ORDER BY Id DESC LIMIT 4").Scan(&TransactionsArray)
	if !strings.Contains(MiningResponse.AddTxStatus, "mining") && len(TransactionsArray) >= 4 {
		if len(TransactionsArray) >= 4 {
			log.Println("11")
			for _, addr := range addresses {
				_, errNode = client.R().SetSuccessResult(&MiningResponse).
					Get(fmt.Sprintf("http://%s/check", strings.Trim(addr.String(), "\"")))
				if errNode != nil {
					if strings.Contains(errNode.Error(), "No connection could be made because the target machine actively refused it.") {
						fmt.Println("aaaaaaaaaaaaa")
						continue
					}
				}
				log.Println("addr", addr)
				for i := 0; i < 4; i++ {
					log.Println("tra", i)
					resp, errReq := client.R().
						SetBody(DeserializeTX(&TransactionsArray[i].Transactions)).
						SetSuccessResult(&txStatus).
						Post(fmt.Sprintf("http://%s/addtx", strings.Trim(addr.String(), "\"")))
					if errReq != nil && !strings.Contains(errReq.Error(), "No connection could be made because the target machine actively refused it.") {
						return "", errReq
					}
					if errReq == nil {
						if resp.Body == nil {
							continue
						}
					}
				}
				for _, v := range TransactionsArray {
					fmt.Println("DELETE FROM TransactionQueue WHERE Id = $1", v.Id)
					db.Exec("DELETE FROM TransactionQueue WHERE Id = $1", v.Id)
				}
				TransactionsArray = TransactionsArray[:0]
			}
		}
	} else {
		log.Println("=================================================================================")
		tx, errSerialize := SerializeTX(&transactionToNet)
		if errSerialize != nil {
			return "", errSerialize
		}
		rand.New(rand.NewSource(time.Now().Unix()))
		t := rand.Intn(10000)
		log.Println(t)
		db.Exec("INSERT INTO TransactionQueue (Id, Transactions) VALUES ($1, $2)", t, tx)
	}

	return txStatus.TransactionStatus, nil
}

func SchedulerTask(client *req.Client, addresses *fastjson.Value) {
	_, err := client.R().SetSuccessResult(&MiningResponse).
		Get(fmt.Sprintf("http://%s/check", strings.Trim(addresses.String(), "\"")))
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadAddresses() ([]*fastjson.Value, error) {
	file, err := os.ReadFile("LowConf/addr.json")
	if err != nil {
		return nil, err
	}
	var p fastjson.Parser
	v, err := p.Parse(string(file))
	if err != nil {
		return nil, err
	}
	return v.GetArray("addresses"), nil
}

//addresses, err := readAddresses()
//if err != nil {
//return "", err
//}

//client := req.C().DevMode()
//_, err := client.R().
//SetBody(&block).
//SetSuccessResult(&result).
//Post(fmt.Sprintf("http://%s/addblock", strings.Trim(goAddr, "\"")))

func SerializeTX(tx *Transport.TransactionHelp) (string, error) {
	jsonData, err := json.MarshalIndent(*tx, "", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func DeserializeTX(data *string) *Transport.TransactionHelp {
	var tx Transport.TransactionHelp
	err := json.Unmarshal([]byte(*data), &tx)
	if err != nil {
		return nil
	}
	return &tx
}
