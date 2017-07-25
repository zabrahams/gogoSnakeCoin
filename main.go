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

	fmt.Print(genBlock.String())
}
