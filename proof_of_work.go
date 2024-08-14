package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

var (
	targetBits = 16
	maxNonce   = 2147483647
)

type ProofOfWork struct {
	block  *Block
	target big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	targetProof := target.Lsh(target, uint(256-targetBits))

	return &ProofOfWork{
		block:  block,
		target: *targetProof,
	}
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.LittleEndian, num)
	if err != nil {
		log.Panic("error while trying to convert int to hex presentation")
	}

	fmt.Println(strconv.FormatInt(num, 16))
	fmt.Println(buff.Bytes())
	fmt.Print("\n\n")

	return buff.Bytes()
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.Data[:],
		pow.block.PrevBlockHash[:],
		IntToHex(int64(pow.block.Timestamp)),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	}, []byte{})

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNonce {
		hash = sha256.Sum256(pow.prepareData(nonce))
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(&pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	var hashInt big.Int
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(&pow.target) == -1 {
		return true
	}

	return false
}
