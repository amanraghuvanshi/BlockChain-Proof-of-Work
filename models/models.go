package models

type Block struct {
}

type Book struct {
}

type BookCheckout struct {
}

type BlockChain struct {
	// slice of multiple blocks, list of blocks
	blocks []*Block
}
