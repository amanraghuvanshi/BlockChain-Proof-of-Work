package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"encoding/json"

	"github.com/amanraghuvanshi/golang-blockchain/models"
)

func (b *models.Block) GenerateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := string(b.Position) + b.TimeStamp + string(bytes) + b.PrevHash

	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevBlock *models.Block, checkoutItem models.BookCheckout) *models.Block {
	block := &models.Block{}
	block.Position = prevBlock.Position + 1
	block.TimeStamp = time.Now().String()
	block.PrevHash = prevBlock.Hash

	// generation of hash
	block.GenerateHash()

	return block
}
