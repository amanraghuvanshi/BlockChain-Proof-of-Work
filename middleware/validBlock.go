package middleware

import (
	"github.com/amanraghuvanshi/golang-blockchain/models"
)

func (b *models.Block) validateHash(hash string) bool {
	b.GenerateHash()
	if b.Hash != hash {
		return false
	}
	return true
}

func ValidBlock(block, prevBlock *models.Block) bool {
	if prevBlock.Hash != block.Hash {
		return false
	}

	if !block.validateHash(block.Hash) {
		return false
	}

	if prevBlock.Position+1 != block.Position {
		return false
	}
	return true
}
