package crawler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestGetTransactionByBlockId(t *testing.T) {
	blockHash := GetBlockHashByBlockId(704328)
	actualBlockHash := "00000000000000000000c2c24d96c099dd51bbfae038022ec6340b077ac3b2e0"
	assert.Equal(t, blockHash, actualBlockHash, "failed to block hash ")
}

func TestGetTransactionHashByBlockHash(t *testing.T) {
	blockHash := "00000000000000000000c2c24d96c099dd51bbfae038022ec6340b077ac3b2e0"
	GetTransactionHashByBlockHash(blockHash)
}

func TestGetTransactionHashByBlockId(t *testing.T) {
	blockId := 704338
	f, err := os.Create(fmt.Sprintf("D://%d.txt", blockId))
	if err != nil {
		log.Fatalf("can't create file: %v\n", err)
	}
	defer f.Close()
	txs := GetTransactionHashByBlockId(blockId)
	for _, tx := range txs {
		f.WriteString(fmt.Sprintf("%s\n", tx))
	}
}
