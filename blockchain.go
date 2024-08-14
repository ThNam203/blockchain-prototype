package main

type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{blocks: []*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) AddBlock(data []byte) {
	lastestBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, lastestBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
