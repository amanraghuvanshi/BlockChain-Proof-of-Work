package controllers

import "github.com/amanraghuvanshi/golang-blockchain/models"

func GenesisBlock() *models.Block {
	return CreateBlock(&models.Block{}, models.BookCheckout{IsGenesis: true})
}

func NewBlockChain() *models.BlockChain {
	return &models.BlockChain{[]*models.Block{GenesisBlock()}}
}
