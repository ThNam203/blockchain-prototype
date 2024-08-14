package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	blockchain := NewBlockchain()
	blockchain.AddBlock([]byte("Block 1 - 1 Bitcoin"))
	blockchain.AddBlock([]byte("Block 2 - 1 Bitcoin"))

	for i, block := range blockchain.blocks {
		pow := NewProofOfWork(block)
		assert.True(t, pow.Validate())

		if i == 0 {
			assert.Equal(t, block.Data, []byte("Genesis block"))
			continue
		}
		assert.Equal(t, block.Data, []byte(fmt.Sprintf("Block %d - 1 Bitcoin", i)))
		t.Log(fmt.Sprintf("Block %d: %+v", i, block))
	}
}
