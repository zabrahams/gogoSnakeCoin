package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}

func (t *Transaction) String() string {
	return fmt.Sprintf("From: %s\nTo: %s\nAmount: %d\n", t.From, t.To, t.Amount)
}

func (t *Transaction) Validate() error {
	if t.From == "" || t.To == "" || t.Amount == 0 {
		return errors.New("error invalid transaction")
	}

	return nil
}
