package models

type Block struct {
}

type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_data"`
	ISBN          string `json:"ISBN"`
}

type BookCheckout struct {
}

type BlockChain struct {
	// slice of multiple blocks, list of blocks
	blocks []*Block
}
