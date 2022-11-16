package main

import (
	"fmt"
	"log"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddress := "my_blockchain_address"

	blockChain := NewBlockchain(myBlockchainAddress)
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my_blockchain_address: %.1f\n", blockChain.CalculateTotalAmount(myBlockchainAddress))
	fmt.Printf("A: %.1f\n", blockChain.CalculateTotalAmount("A"))
	fmt.Printf("B: %.1f\n", blockChain.CalculateTotalAmount("B"))
	fmt.Printf("C: %.1f\n", blockChain.CalculateTotalAmount("C"))
	fmt.Printf("D: %.1f\n", blockChain.CalculateTotalAmount("D"))
	fmt.Printf("X: %.1f\n", blockChain.CalculateTotalAmount("X"))
	fmt.Printf("Y: %.1f\n", blockChain.CalculateTotalAmount("Y"))
}
