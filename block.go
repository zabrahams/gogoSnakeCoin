package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index        int64
	Timestamp    time.Time
	Data         []byte
	PreviousHash []byte
	Hash         []byte
}

func (b Block) MakeHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte{byte(b.Index)}); err != nil {
		return nil, err
	}
	if _, err := h.Write([]byte(b.Timestamp.Format(time.RFC3339))); err != nil {
		return nil, err
	}
	if _, err := h.Write(b.Data); err != nil {
		return nil, err
	}
	if _, err := h.Write(b.PreviousHash); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func (b Block) String() string {
	return fmt.Sprintf("Index: %d\nTimestamp: %v\nData: %s\nPreviousHash: %X\nHash: %X\n", b.Index, b.Timestamp, b.Data, b.PreviousHash, b.Hash)
}
func CreateGenesisBlock() (*Block, error) {
	genBlock := &Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         []byte("Genesis Block"),
		PreviousHash: []byte("0"),
	}
	h, err := genBlock.MakeHash()
	if err != nil {
		return nil, err
	}
	genBlock.Hash = h
	return genBlock, nil
}
