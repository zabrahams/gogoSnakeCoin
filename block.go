package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index        int64     `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	Data         BlockData `json:"data"`
	PreviousHash []byte    `json:"previous_hash"`
	Hash         []byte    `json:"hash"`
}

type BlockData struct {
	ProofOfWork  int           `json:"proof_of_work"`
	Transactions []Transaction `json:"transactions"`
}

func (bd BlockData) String() string {
	return fmt.Sprintf("Proof: %d\nTransactions: %+v", bd.ProofOfWork, bd.Transactions)
}

func (b Block) MakeHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte{byte(b.Index)}); err != nil {
		return nil, err
	}
	if _, err := h.Write([]byte(b.Timestamp.Format(time.RFC3339))); err != nil {
		return nil, err
	}
	if _, err := h.Write([]byte{byte(b.Data.ProofOfWork)}); err != nil {
		return nil, err
	}
	for _, txion := range b.Data.Transactions {
		jsonTxion := []byte(txion.String())
		if _, err := h.Write(jsonTxion); err != nil {
			return nil, err
		}
	}
	if _, err := h.Write(b.PreviousHash); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func (b Block) String() string {
	return fmt.Sprintf("Index: %d\nTimestamp: %v\nData: %s\nPreviousHash: %X\nHash: %X\n", b.Index, b.Timestamp, b.Data, b.PreviousHash, b.Hash)
}
func CreateGenesisBlock() (Block, error) {
	genBlock := Block{
		Index:     0,
		Timestamp: time.Now(),
		Data: BlockData{
			ProofOfWork:  1,
			Transactions: []Transaction{},
		},
		PreviousHash: []byte("0"),
	}
	h, err := genBlock.MakeHash()
	if err != nil {
		return Block{}, err
	}
	genBlock.Hash = h
	return genBlock, nil
}

func (b Block) NextBlock(proof int, txions []Transaction) (Block, error) {
	newBlock := Block{
		Index:     b.Index + 1,
		Timestamp: time.Now(),
		Data: BlockData{
			ProofOfWork:  proof,
			Transactions: txions,
		},
		PreviousHash: b.Hash,
	}
	h, err := newBlock.MakeHash()
	if err != nil {
		return Block{}, err
	}
	newBlock.Hash = h
	return newBlock, nil
}
