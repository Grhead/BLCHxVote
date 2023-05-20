package Blockchain

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sort"
)

func LastHashCompare(master string) (string, error) {
	db, err := gorm.Open(sqlite.Open("Database/CompareDb.db"), &gorm.Config{})
	if err != nil {
		return "", err
	}
	var hash string
	var chain []*Chain
	var blocks []*Block
	db.Find(&chain)
	for _, v := range chain {
		deserializedBlock, err := DeserializeBlock(v.Block)
		if err != nil {
			return "", err
		}
		blocks = append(blocks, deserializedBlock)
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

func AddBlockCompare(block *Block) error {
	//db, err := gorm.Open(sqlite.Open("Database/CompareDb.db"), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("DatabaseTest/CompareDb.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	serializedBlock, err := SerializeBlock(block)
	if err != nil {
		return err
	}
	errInsert := db.Exec("INSERT INTO Chains (Id, Hash, Block) VALUES ($1, $2, $3)",
		uuid.NewString(),
		block.CurrHash,
		serializedBlock,
	)
	if errInsert.Error != nil {
		return errInsert.Error
	}
	return nil
}
