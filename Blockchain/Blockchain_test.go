package Blockchain

//func TestAddBlock(t *testing.T) {
//	db, err_open := gorm.Open(sqlite.Open("Database/NodeDb.db"), &gorm.Config{})
//	assert.NoError(t, err_open)
//	chain := &BlockChain{
//		DB: db,
//	}
//	curTime, err_time := GetTime()
//	assert.NoError(t, err_time)
//	var block *Block
//	block.TimeStamp = curTime.AsTime().String()
//	block.ChainMaster = "Test"
//	assert.NoError(t, chain.AddBlock(block))
//}
