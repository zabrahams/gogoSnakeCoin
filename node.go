package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Node struct {
	Address      string
	Blockchain   []Block
	Transactions []Transaction
	NodeURLs     []string
}

func (n *Node) Mine() error {
	lastBlock := n.Blockchain[len(n.Blockchain)-1]
	lastProof := lastBlock.Data.ProofOfWork

	proof := ProofOfWork(lastProof)

	n.Transactions = append(n.Transactions, Transaction{
		From:   "network",
		To:     n.Address,
		Amount: 1,
	})
	newBlock, err := lastBlock.NextBlock(proof, n.Transactions)
	if err != nil {
		return err
	}

	n.Transactions = []Transaction{}
	n.Blockchain = append(n.Blockchain, newBlock)
	return nil
}

func (n *Node) GetChains() [][]Block {
	var blockchains [][]Block
	for _, url := range n.NodeURLs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("error fetching blockchain from %s - %+v", url, err)
			continue
		}
		blockchain := []Block{}
		defer resp.Body.Close()
		jsonBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error reading response from %s - %+v", url, err)
			continue
		}
		if err := json.Unmarshal(jsonBody, blockchain); err != nil {
			fmt.Printf("error parsing blockhain json from %s - %+v", url, err)
			continue
		}

		blockchains = append(blockchains, blockchain)
	}

	return blockchains
}

func (n *Node) Consensus() {
	longestChain := n.Blockchain
	blockchains := n.GetChains()
	for _, blockchain := range blockchains {
		if len(blockchain) > len(longestChain) {
			longestChain = blockchain
		}
	}

	n.Blockchain = longestChain
}

func ProofOfWork(lastProof int) int {
	inc := lastProof + 1

	for !(inc%9 == 0 && inc%lastProof == 0) {
		inc = inc + 1
	}

	return inc
}
