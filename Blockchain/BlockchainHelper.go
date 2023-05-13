package Blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// SerializeBlock TODO Rewrite
func SerializeBlock(block *Block) string {
	jsonData, err := json.MarshalIndent(*block, "", "\t")
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// DeserializeBlock TODO Rewrite
func DeserializeBlock(data string) *Block {
	var block Block
	err := json.Unmarshal([]byte(data), &block)
	if err != nil {
		return nil
	}
	return &block
}

// SerializeTX TODO Rewrite
func SerializeTX(tx *Transaction) string {
	jsonData, err := json.MarshalIndent(*tx, "", "\t")
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// DeserializeTX TODO Rewrite
func DeserializeTX(data string) *Transaction {
	var tx Transaction
	err := json.Unmarshal([]byte(data), &tx)
	if err != nil {
		return nil
	}
	return &tx
}

func GetTime() (*timestamp.Timestamp, error) {
	TimeUrl := viper.GetString("TIME_URL")
	resp, err := http.Get(TimeUrl)
	if err != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)
		return timestamppb.New(time.Time{}), err
	}
	//TODO what is log fatal
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return timestamppb.New(time.Time{}), err
	}
	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return timestamppb.New(time.Time{}), err
	}
	parsedTime, err := time.Parse("2006-01-02T15:04:05", string(v.GetStringBytes("dateTime")))
	if err != nil {
		return timestamppb.New(time.Time{}), err
	}
	return timestamppb.New(parsedTime), nil

}

func SetHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return string(hash[:])
}

func ToBytes(data uint64) []byte {
	var buf = new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func GenerateRandomBytes(max uint64) ([]byte, error) {
	var slice = make([]byte, max)
	_, err := rand.Read(slice)
	if err != nil {
		return nil, err
	}
	return slice, nil
}

func HashSum(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash[:])
}

//func ProofOfWork(blockHash []byte, difficulty uint8, ch chan bool) uint64 {
//	DEBUG := true
//	var (
//		Target  = big.NewInt(1)
//		intHash = big.NewInt(1)
//		nonce   = uint64(rand.Intn(math.MaxUint32))
//		hash    []byte
//	)
//	Target.Lsh(Target, 256-uint(difficulty))
//	for nonce < math.MaxUint64 {
//		select {
//		case <-ch:
//			if DEBUG {
//				fmt.Println()
//			}
//			return nonce
//		default:
//			hash = []byte(HashSum(bytes.Join(
//				[][]byte{
//					blockHash,
//					ToBytes(nonce),
//				},
//				[]byte{},
//			)))
//			if DEBUG {
//				fmt.Printf("\rMining: %s", Base64Encode(hash))
//			}
//			intHash.SetBytes(hash)
//			if intHash.Cmp(Target) == -1 {
//				if DEBUG {
//					fmt.Println()
//				}
//				return nonce
//			}
//			nonce++
//		}
//	}
//	return nonce
//}
