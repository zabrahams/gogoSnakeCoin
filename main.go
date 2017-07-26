package main

import (
	"fmt"
	"log"
)

func main() {
	genBlock, err := CreateGenesisBlock()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	node := Node{
		Address:    "testaddress",
		Blockchain: []Block{genBlock},
	}
	serv := &Server{
		Node: node,
	}
	serv.Start("8080")
}

func genSimpleChain() {
	genBlock, err := CreateGenesisBlock()
	if err != nil {
		log.Fatalf("error creating genesis block: %+v", err)
	}

	fmt.Println("Genesis Block Created:")
	blockchain := []Block{genBlock}
	for i := 0; i <= 20; i += 1 {
		newBlock, err := blockchain[i].NextBlock(1, []Transaction{})
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
