package Basic

import (
	"VOX2/Blockchain"
	"fmt"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestViewCandidate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../Database/ContractDB.db"), &gorm.Config{})
	if err != nil {
		return
	}
	var Candidates []*Blockchain.ElectionSubjects
	db.Find(&Candidates)
	assert.Equal(t, Candidates, nil)
}

func TestNewCandidate(t *testing.T) {
	viper.SetConfigFile("../LowConf/config.env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
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
		return
	}
	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return
	}
	hash := v.GetStringBytes("dateTime")
	db, err := gorm.Open(sqlite.Open("../Database/ContractDB.db"), &gorm.Config{})
	if err != nil {
		return
	}
	tempUUID := uuid.New()
	tempKey := Blockchain.HashSum(string(hash))
	db.Exec("INSERT INTO ElectionSubjects (Id, PublicKey,Description, VotingAffiliation) VALUES ($1, $2, $3, $4)",
		tempUUID,
		tempKey,
		"TEST1",
		"Master")
	candidate := Blockchain.ElectionSubjects{
		Uuid:              tempUUID,
		PublicKey:         tempKey,
		Description:       "TEST1",
		VotingAffiliation: "Master",
	}
	assert.Equal(t, candidate, Blockchain.ElectionSubjects{
		Uuid:              tempUUID,
		PublicKey:         tempKey,
		Description:       "TEST1",
		VotingAffiliation: "Master",
	})
}
