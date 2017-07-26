package main

type Node struct {
	Address      string
	Blockchain   []Block
	Transactions []Transaction
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

func ProofOfWork(lastProof int) int {
	inc := lastProof + 1

	for !(inc%9 == 0 && inc%lastProof == 0) {
		inc = inc + 1
	}

	return inc
}
