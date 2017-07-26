package main

import (
	"encoding/json"
	"errors"
)

type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}

func (t *Transaction) String() string {
	jsonTx, err := json.Marshal(&t)
	if err != nil {
		return "bad txion"
	}
	return string(jsonTx)
}

func (t *Transaction) Validate() error {
	if t.From == "" || t.To == "" || t.Amount == 0 {
		return errors.New("error invalid transaction")
	}

	return nil
}
