package main

import (
	"log"
	"net/http"

	"github.com/amanraghuvanshi/golang-blockchain/controllers"
	"github.com/amanraghuvanshi/golang-blockchain/models"
	"github.com/gorilla/mux"
)

var BlockChain *models.BlockChain

func (bc *BlockChain) AddBlock(data models.BookCheckout) {
	prevBlock := bc.block[len(bc.block)-1]
	// creation of block
	block := controllers.CreateBlock(prevBlock, data)
	//check the validity
	if validBlock(block, prevBlock) {
		bc.block = append(bc.block, block)
	}
}

func main() {
	// router
	r := mux.NewRouter()
	// Get you chain
	r.HandleFunc("/", getTheChain).Methods("GET")
	// second route
	r.HandleFunc("/", controllers.WriteBlock()).Methods("POST")
	// the first route
	r.HandleFunc("/new", controllers.NewBook()).Methods("POST")

	//confirmation
	log.Println("Listening at port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
