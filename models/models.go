package models

type Block struct {
	PrevHash  string
	Position  int
	Data      BookCheckout
	TimeStamp string
	Hash      string
}

type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_data"`
	ISBN          string `json:"ISBN"`
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	isGenesis    bool   `json:"is_genesis"`
}

type BlockChain struct {
	// slice of multiple blocks, list of blocks
	blocks []*Block
}
