package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     uint64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(int64(b.Timestamp), 10))
	concatenated := bytes.Join([][]byte{b.PrevBlockHash, timestamp, b.Data}, []byte{})
	hashedSum := sha256.Sum256(concatenated)
	b.Hash = hashedSum[:]
}

func NewBlock(data []byte, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     uint64(time.Now().Unix()),
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}

	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock([]byte("Genesis block"), []byte{})
}

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

func main() {

}
