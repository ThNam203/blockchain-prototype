package main

import (
	"time"
)

type Block struct {
	Timestamp     uint64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data []byte, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     uint64(time.Now().Unix()),
		Data:          data,
		PrevBlockHash: prevBlockHash,
		Nonce:         0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock([]byte("Genesis block"), []byte{})
}
