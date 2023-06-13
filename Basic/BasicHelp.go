package Basic

import (
	"VOX2/Transport"
	"github.com/goccy/go-json"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/valyala/fastjson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const AllowedError = "No connection could be made because the target machine actively refused it."

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

func setTime(master string, limit *timestamppb.Timestamp) (*timestamp.Timestamp, error) {
	log.Println("setTime")
	db, err := gorm.Open(sqlite.Open("Database/ContractDB.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	rand.New(rand.NewSource(time.Now().Unix()))
	t := rand.Intn(10000)
	errInsert := db.Exec("INSERT INTO VotingTime (Id, MasterChain, LimitTime) VALUES ($1, $2, $3)", t, master, limit.Seconds)
	if errInsert.Error != nil {
		return nil, errInsert.Error
	}
	return limit, nil
}

func checkTime(master string) (time.Time, string, error) {
	log.Println("checkTime")
	db, err := gorm.Open(sqlite.Open("Database/ContractDB.db"), &gorm.Config{})
	if err != nil {
		return time.Time{}, "", err
	}
	var timeOfMaster string
	db.Raw("SELECT LimitTime FROM VotingTime WHERE MasterChain = $1",
		master).Scan(&timeOfMaster)
	i, err := strconv.Atoi(timeOfMaster)
	if err != nil {
		return time.Time{}, "", err
	}
	parsedTime := time.Unix(int64(i), 0).UTC()
	return parsedTime, timeOfMaster, nil
}

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
