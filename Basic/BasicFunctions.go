package Basic

import (
	"VOX2/Blockchain"
	"VOX2/Transport"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-co-op/gocron"
	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type TransactionsDb struct {
	Id           int
	Transactions string
	Master       string
}
type ElectionsList struct {
	ElectionSubject *Blockchain.ElectionSubjects
	Balance         string
}

func init() {
	viper.SetConfigFile("./LowConf/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

var MiningResponse Transport.CheckHelp

func NewChain(initMaster string, votesCount int64, limit *timestamppb.Timestamp) (*Transport.CreateHelp, error) {
	log.Println("NewChain")
	addresses, err := ReadAddresses()
	if err != nil {
		return nil, err
	}
	initChain := Transport.ChainHelp{
		Master: initMaster,
		Count:  votesCount,
	}

	var creation Transport.CreateHelp
	client := req.C().DevMode()
	for _, addr := range addresses {
		resp, errReq := client.R().
			SetBody(&initChain).
			SetSuccessResult(&creation).
			Post(fmt.Sprintf("http://%s/newchain", strings.Trim(addr.String(), "\"")))
		if errReq != nil && !strings.Contains(errReq.Error(), AllowedError) {
			return nil, errReq
		}
		if errReq == nil {
			if resp.Body == nil {
				continue
			}
		}
	}
	_, err = setTime(initMaster, limit)
	if err != nil {
		return nil, err
	}
	return &creation, nil
}

func CallCreateVoters(voter string, master string) ([]*Blockchain.User, []string, error) {
	log.Println("CallCreateVoters")
	DbConf := viper.GetString("DCS")
	db, err := gorm.Open(mysql.Open(DbConf), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	rand.New(rand.NewSource(time.Now().Unix()))
	var resultItems []*Blockchain.User
	var resultItemsPass []string
	_, errVoter := strconv.Atoi(voter)
	getTime, errGetTime := Blockchain.GetTime()
	if errGetTime != nil {
		return nil, nil, errGetTime
	}
	checkTimeVar, _, err := checkTime(master)
	if err != nil {
		return nil, nil, err
	}
	if getTime.AsTime().After(checkTimeVar) {
		log.Println("time expired")
		return nil, nil, errors.New("time expired")
	}
	if errVoter != nil {
		PseudoIdentity := Blockchain.HashSum(voter)[:16]
		_, errDormantUser := Blockchain.NewDormantUser(PseudoIdentity, master)
		if errDormantUser != nil {
			return nil, nil, errDormantUser
		}
		_, errPublicKeyItem := Blockchain.NewPublicKeyItem(master)
		if errPublicKeyItem != nil {
			return nil, nil, errPublicKeyItem
		}
		if errPublicKeyItem != nil {
			return nil, nil, errPublicKeyItem
		}
		DormantUserRows, errRaw := db.Raw("SELECT PersonIdentifier FROM RelationPatterns WHERE Master = ?", master).Rows()
		if errRaw != nil {
			return nil, nil, errRaw
		}
		rows, errRows := db.Raw("SELECT * FROM PublicKeySets WHERE VotingAffiliation = ?", master).Rows()
		defer func(rows *sql.Rows) {
			errClose := rows.Close()
			if errClose != nil {
				log.Fatalln(errClose)
			}
		}(rows)
		if errRows != nil {
			return nil, nil, errRows
		}
		for rows.Next() {
			var tempItem Blockchain.User
			errScan := rows.Scan(&tempItem.Id, &tempItem.PublicKey, &tempItem.IsUsed, &tempItem.VotingAffiliation)
			if errScan != nil {
				return nil, nil, errScan
			}
			resultItems = append(resultItems, &tempItem)
		}
		for DormantUserRows.Next() {
			var tempItem string
			errScan := DormantUserRows.Scan(&tempItem)
			if errScan != nil {
				return nil, nil, errScan
			}
			aes, errDecryptAES := Blockchain.DecryptAES([]byte(master), tempItem)
			if errDecryptAES != nil {
				return nil, nil, errDecryptAES
			}
			if errScan != nil {
				return nil, nil, errScan
			}
			resultItemsPass = append(resultItemsPass, aes)
		}
	} else {
		max, errAtom := strconv.Atoi(fmt.Sprintf("%v", voter))
		var PseudoIdentity string
		if errAtom != nil {
			return nil, nil, errAtom
		}
		for i := 0; i < max; i++ {
			rand.New(rand.NewSource(time.Now().Unix()))
			t := rand.Int63()
			PseudoIdentity = Blockchain.HashSum(fmt.Sprintf("%x", getTime.AsTime().Unix()+t))[:16]
			_, errDormant := Blockchain.NewDormantUser(PseudoIdentity, master)
			if errDormant != nil {
				return nil, nil, errDormant
			}
			_, errNewPublicKey := Blockchain.NewPublicKeyItem(master)
			if errNewPublicKey != nil {
				return nil, nil, errNewPublicKey
			}
		}
		DormantUserRows, errRaw := db.Raw("SELECT PersonIdentifier FROM RelationPatterns WHERE Master = ?", master).Rows()
		if errRaw != nil {
			return nil, nil, errRaw
		}
		rows, errRows := db.Raw("SELECT * FROM PublicKeySets WHERE VotingAffiliation = ?", master).Rows()
		defer func(rows *sql.Rows) {
			errClose := rows.Close()
			if errClose != nil {
				log.Fatalln(errClose)
			}
		}(rows)
		if errRows != nil {
			return nil, nil, errRows
		}
		for rows.Next() {
			var tempItem Blockchain.User
			errScan := rows.Scan(&tempItem.Id, &tempItem.PublicKey, &tempItem.IsUsed, &tempItem.VotingAffiliation)
			if errScan != nil {
				return nil, nil, errScan
			}
			resultItems = append(resultItems, &tempItem)
		}
		for DormantUserRows.Next() {
			var tempItem string
			errScan := DormantUserRows.Scan(&tempItem)
			if errScan != nil {
				return nil, nil, errScan
			}
			aes, errDecryptAES := Blockchain.DecryptAES([]byte(master), tempItem)
			if errDecryptAES != nil {
				return nil, nil, errDecryptAES
			}
			if errScan != nil {
				return nil, nil, errScan
			}
			resultItemsPass = append(resultItemsPass, aes)
		}
	}
	return resultItems, resultItemsPass, nil
}
func ViewEffectivity(master string) (int, error) {
	var resultItemsPass []string
	DbConf := viper.GetString("DCS")
	db, err := gorm.Open(mysql.Open(DbConf), &gorm.Config{})
	if err != nil {
		return 0, err
	}
	DormantUserRows, errRaw := db.Raw("SELECT PersonIdentifier FROM RelationPatterns WHERE Master = ?", master).Rows()
	if errRaw != nil {
		return 0, errRaw
	}
	rows, errRows := db.Raw("SELECT * FROM PublicKeySets WHERE VotingAffiliation = ?", master).Rows()
	defer func(rows *sql.Rows) {
		errClose := rows.Close()
		if errClose != nil {
			log.Fatalln(errClose)
		}
	}(rows)
	if errRows != nil {
		return 0, errRows
	}
	for DormantUserRows.Next() {
		var tempItem string
		errScan := DormantUserRows.Scan(&tempItem)
		if errScan != nil {
			return 0, errScan
		}
		aes, errDecryptAES := Blockchain.DecryptAES([]byte(master), tempItem)
		if errDecryptAES != nil {
			return 0, errDecryptAES
		}
		if errScan != nil {
			return 0, errScan
		}
		resultItemsPass = append(resultItemsPass, aes)
	}
	return len(resultItemsPass), nil
}

func CallViewCandidates(master string) ([]*Blockchain.ElectionSubjects, error) {
	log.Println("CallViewCandidates", " Master =", master)
	DbConf := viper.GetString("DCS")
	db, err := gorm.Open(mysql.Open(DbConf), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var Candidates []*Blockchain.ElectionSubjects
	db.Table("ElectionSubjects").Where("VotingAffiliation = ?", master).Find(&Candidates)
	fmt.Println(Candidates)
	return Candidates, nil
}

func CallNewCandidate(description string, affiliation string) (*Blockchain.ElectionSubjects, error) {
	log.Println("CallNewCandidate")
	gettingTime, err := Blockchain.GetTime()
	if err != nil {
		return nil, err
	}
	log.Println(gettingTime)
	checkTimeVar, _, err := checkTime(affiliation)
	if err != nil {
		return nil, err
	}
	log.Println(checkTimeVar)
	if gettingTime.AsTime().After(checkTimeVar) {
		log.Println("time expired")
		return nil, errors.New("time expired")
	}
	list, err := CallViewCandidates(affiliation)
	for _, v := range list {
		if description == v.Description {
			return nil, errors.New("election object already exist")
		}
	}
	candidate, err := Blockchain.NewCandidate(description, affiliation)
	if err != nil {
		return nil, err
	}
	log.Println(candidate)
	return candidate, nil
}

func WinnersList(master string) ([]*ElectionsList, error) {
	log.Println("WinnersList")
	DbConf := viper.GetString("DCS")
	db, errDb := gorm.Open(mysql.Open(DbConf), &gorm.Config{})
	if errDb != nil {
		return nil, errDb
	}
	var GetElections []*Blockchain.ElectionSubjects
	var ResultList []*ElectionsList
	db.Table("ElectionSubjects").Where("VotingAffiliation = ?", master).Find(&GetElections)
	if len(GetElections) < 1 {
		return nil, errors.New("empty list")
	}
	for _, v := range GetElections {
		ElectionBalance, err := getBalance(v.PublicKey)
		if err != nil {
			return nil, err
		}
		ElectionItem := ElectionsList{
			ElectionSubject: v,
			Balance:         ElectionBalance.Balance,
		}
		ResultList = append(ResultList, &ElectionItem)
	}
	return ResultList, nil
}

func SoloWinner(master string) (*ElectionsList, error) {
	log.Println("SoloWinner")
	GetElections, err := WinnersList(master)
	if err != nil {
		return nil, err
	}
	sort.Slice(GetElections, func(i, j int) bool {
		GetBalanceI, errBalanceI := strconv.Atoi(GetElections[j].Balance)
		if errBalanceI != nil {
			log.Fatalln(errBalanceI)
		}
		GetBalanceJ, errBalanceJ := strconv.Atoi(GetElections[j].Balance)
		if errBalanceJ != nil {
			log.Fatalln(errBalanceJ)
		}
		return GetBalanceI < GetBalanceJ
	})
	return GetElections[len(GetElections)-1], nil
}

func getBalance(userAddress string) (*Transport.BalanceHelp, error) {
	log.Println("getBalance")
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
		if errReq != nil && !strings.Contains(errReq.Error(), AllowedError) {
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
	log.Println("ChainSize")
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
	log.Println("GetPartOfChain")
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
	log.Println("GetFullChain")
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

func AcceptNewUser(Pass string, salt string, PublicKey string) (string, error) {
	log.Println("AcceptNewUser")
	master, err := Blockchain.GetVotingAffiliation(PublicKey)
	if err != nil {
		return "", err
	}
	gettingTime, err := Blockchain.GetTime()
	if err != nil {
		return "", err
	}
	checkTimeVar, _, err := checkTime(master)
	if err != nil {
		return "", err
	}
	if gettingTime.AsTime().After(checkTimeVar) {
		return "", errors.New("time expired")
	}
	private, err := Blockchain.RegisterGeneratePrivate(Pass, salt, PublicKey)
	if err != nil {
		return "", err
	}
	_, errTransfer := transfer(PublicKey, master, 1)
	if errTransfer != nil {
		return "", errTransfer
	}
	return private, nil
}

func AcceptLoadUser(PublicK string, PrivateK string) (*Blockchain.User, error) {
	log.Println("AcceptLoadUser")
	DbConf := viper.GetString("DCS")
	db, err := gorm.Open(mysql.Open(DbConf), &gorm.Config{})
	master, err := Blockchain.GetVotingAffiliation(PublicK)
	if err != nil {
		return nil, err
	}
	var timeOfMaster string
	db.Raw("SELECT LimitTime FROM VotingTime WHERE MasterChain = ?",
		master).Scan(&timeOfMaster)
	if timeOfMaster == "" {
		return nil, errors.New("invalid master voting")
	}
	gettingTime, err := Blockchain.GetTime()
	if err != nil {
		return nil, err
	}
	checkTimeVar, _, err := checkTime(master)
	if err != nil {
		return nil, err
	}
	if gettingTime.AsTime().After(checkTimeVar) {
		return nil, errors.New("time expired")
	}
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
	bal, err := getBalance(UserPublic.Address())
	if err != nil {
		return nil, err
	}
	if bal.Balance == "" {
		return nil, errors.New("zero balance")
	}
	log.Println(UserPublic)
	return UserPublic, nil
}

func Vote(receiver string, sender string, master string, num int64) (string, error) {
	log.Println("Vote")
	DbConf := viper.GetString("DCS")
	db, errDb := gorm.Open(mysql.Open(DbConf), &gorm.Config{})
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
		Master: sender,
	}
	var errNode error
	for _, addr := range addresses {
		_, errNode = client.R().SetSuccessResult(&MiningResponse).
			Get(fmt.Sprintf("http://%s/check", strings.Trim(addr.String(), "\"")))
		if errNode != nil {
			if strings.Contains(errNode.Error(), AllowedError) {
				continue
			}
		}
		resp, errReq := client.R().
			SetBody(&RequestData).
			SetSuccessResult(&lastHash).
			Post(fmt.Sprintf("http://%s/getlasthash", strings.Trim(addr.String(), "\"")))
		if errReq != nil && !strings.Contains(errReq.Error(), AllowedError) {
			return "", errReq
		}
		if errReq == nil {
			if resp.Body == nil {
				continue
			}
		}
		balance, errBalance := getBalance(sender)
		if errBalance != nil {
			return "", errBalance
		}
		chainBalance, errConversion := strconv.Atoi(balance.Balance)
		if errConversion != nil {
			return "", errConversion
		}
		if int64(chainBalance) < num {
			return "", errors.New("not enough user founds")
		}
		publicSender, errLoad := Blockchain.LoadToEnterAlreadyUserPublic(sender)
		if errLoad != nil {
			return "", errLoad
		}
		publicReceiver, errLoad := Blockchain.GetCandidate(receiver)
		publicObject := &Transport.ObjectHelp{
			PublicKey:         publicReceiver.PublicKey,
			VotingAffiliation: publicReceiver.VotingAffiliation,
		}
		transactionToNet = Transport.TransactionHelp{
			Master:   master,
			Receiver: publicObject,
			Count:    num,
			Sender:   publicSender,
		}
	}
	tx, errSerialize := SerializeTX(&transactionToNet)
	if errSerialize != nil {
		log.Fatalln(errSerialize)
	}
	db.Exec("INSERT INTO TransactionQueue (Transactions, Master) VALUES (?, ?)", tx, master)
	addTransactionToNet(db, client, addresses, txStatus, master)

	return txStatus.TransactionStatus, nil
}

func transfer(receiver string, master string, num int64) (string, error) {
	log.Println("transfer")
	DbConf := viper.GetString("DCS")
	db, errDb := gorm.Open(mysql.Open(DbConf), &gorm.Config{})
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
	var errNode error
	RequestData := Transport.MasterHelp{
		Master: master,
	}
	for _, addr := range addresses {
		_, errNode = client.R().SetSuccessResult(&MiningResponse).
			Get(fmt.Sprintf("http://%s/check", strings.Trim(addr.String(), "\"")))
		if errNode != nil {
			if strings.Contains(errNode.Error(), AllowedError) {
				continue
			}
		}
		resp, errReq := client.R().
			SetBody(&RequestData).
			SetSuccessResult(&lastHash).
			Post(fmt.Sprintf("http://%s/getlasthash", strings.Trim(addr.String(), "\"")))
		if errReq != nil && !strings.Contains(errReq.Error(), AllowedError) {
			return "", errReq
		}
		if errReq == nil {
			if resp.Body == nil {
				continue
			}
		}
		balance, errBalance := getBalance(master)
		if errBalance != nil {
			return "", errBalance
		}
		chainBalance, errConversion := strconv.Atoi(balance.Balance)
		if errConversion != nil {
			return "", errConversion
		}
		if int64(chainBalance) < num {
			return "", errors.New("not enough chain founds")
		}
		public, errLoad := Blockchain.GetUserByPublic(receiver)
		publicObject := &Transport.ObjectHelp{
			PublicKey:         public.PublicKey,
			VotingAffiliation: public.VotingAffiliation,
		}
		if errLoad != nil {
			return "", errLoad
		}
		transactionToNet = Transport.TransactionHelp{
			Master:   master,
			Receiver: publicObject,
			Count:    num,
		}
	}
	tx, errSerialize := SerializeTX(&transactionToNet)
	if errSerialize != nil {
		log.Fatalln(errSerialize)
	}
	db.Exec("INSERT INTO TransactionQueue (Transactions, Master) VALUES (?, ?)", tx, master)
	log.Println(tx, master)
	addTransactionToNet(db, client, addresses, txStatus, master)
	return txStatus.TransactionStatus, nil
}

func addTransactionToNet(
	db *gorm.DB,
	client *req.Client,
	addresses []*fastjson.Value,
	txStatus Transport.TransactionResponseHelp,
	master string) {
	var TransactionsArray []TransactionsDb
	db.Raw("SELECT Id, Transactions, Master FROM TransactionQueue WHERE Master = ? ORDER BY Id DESC LIMIT 4", master).Scan(&TransactionsArray)
	if !strings.Contains(MiningResponse.AddTxStatus, "mining") && len(TransactionsArray) >= 4 {
		for _, addr := range addresses {
			_, errNode := client.R().SetSuccessResult(&MiningResponse).
				Get(fmt.Sprintf("http://%s/check", strings.Trim(addr.String(), "\"")))
			if errNode != nil {
				if strings.Contains(errNode.Error(), AllowedError) {
					continue
				}
			}
			for i := 0; i < 4; i++ {
				resp, errReq := client.R().
					SetBody(DeserializeTX(&TransactionsArray[i].Transactions)).
					SetSuccessResult(&txStatus).
					Post(fmt.Sprintf("http://%s/addtx", strings.Trim(addr.String(), "\"")))
				if errReq != nil && !strings.Contains(errReq.Error(), AllowedError) {
					log.Fatalln(errReq)
				}
				if errReq == nil {
					if resp.Body == nil {
						continue
					}
				}
			}
		}
		for _, v := range TransactionsArray {
			db.Exec("DELETE FROM TransactionQueue WHERE Id = ?", v.Id)
		}
	}
}
