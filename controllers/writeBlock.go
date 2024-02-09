package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amanraghuvanshi/golang-blockchain/models"
)

func WriteBlock(w http.ResponseWriter, r http.Request) {
	var checkoutItem models.BookCheckout

	if err := json.NewDecoder(r.Body).Decode(&checkoutItem); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Couldnt write the block: %v", err)
		w.Write([]byte('Couldnt write the block'))
	}
	BlockChain.AddBlock(checkoutItem)
}
