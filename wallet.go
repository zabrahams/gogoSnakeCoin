package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Wallet struct {
	KeyRepo KeyRepo
}

type KeyRepo interface {
	StoreKey(*rsa.PrivateKey) (string, error)
}

func (w *Wallet) GenerateKey() (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", fmt.Errorf("error generating keys: %+v", err)
	}
	publicKeyStr, err := w.KeyRepo.StoreKey(privateKey)
	if err != nil {
		return "", fmt.Errorf("error storing new key: %+v", err)
	}

	return publicKeyStr, nil
}
