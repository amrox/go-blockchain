package main

import (
	"fmt"
	"log"

	"github.com/amrox/go-blockchain/pkg/blockchain"
)

func main() {
	log.Println("blockchain")

	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Andy")
	bc.AddBlock("Send 2 more BTC to Andy")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x", block.Hash)
		fmt.Println()
	}
}
