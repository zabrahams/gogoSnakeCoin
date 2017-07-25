package main

import (
	"fmt"
	"log"
)

func main() {
	genBlock, err := CreateGenesisBlock()
	if err != nil {
		log.Fatalf("error creating genesis block: %+v", err)
	}

	fmt.Println("Genesis Block Created:")
	fmt.Print(genBlock.String())

	blockchain := []*Block{genBlock}
	for i := 0; i <= 20; i += 1 {
		newBlock, err := blockchain[i].NextBlock()
		if err != nil {
			log.Fatalf("uh oh - problem building block chain: %+v", err)
		}
		blockchain = append(blockchain, newBlock)
	}

	for _, block := range blockchain {
		fmt.Print(block.String())
		fmt.Println("-----------------------------------------------------")
	}
}
