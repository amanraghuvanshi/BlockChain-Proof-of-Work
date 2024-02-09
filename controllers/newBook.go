package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/amanraghuvanshi/golang-blockchain/models"
)

func newBook(w http.ResponseWriter, r http.Request) {
	// creation of new book
	var book models.Book

	// the information we get in JSON, we need it in struct format for go, so we convert it over here
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Couldnt create the book: %v", err)
		w.Write([]byte("Couldnt create a new book"))
		return
	}

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishedDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))
}
