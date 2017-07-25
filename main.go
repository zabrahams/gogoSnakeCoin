package main

import (
	"fmt"
	"log"
)

func main() {
	serv := &Server{}
	serv.Start("8080")
}

func genSimpleChain() {
	genBlock, err := CreateGenesisBlock()
	if err != nil {
		log.Fatalf("error creating genesis block: %+v", err)
	}

	fmt.Println("Genesis Block Created:")
	blockchain := []*Block{genBlock}
	for i := 0; i <= 20; i += 1 {
		newBlock, err := blockchain[i].NextBlock()
		if err != nil {
			log.Fatalf("uh oh - problem building block chain: %+v", err)
		}
		fmt.Printf("Block #%d Added to the chain.\n", newBlock.Index)
		blockchain = append(blockchain, newBlock)
	}

	for _, block := range blockchain {
		fmt.Print(block.String())
		fmt.Println("-----------------------------------------------------")
	}
}
