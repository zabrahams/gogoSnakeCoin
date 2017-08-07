package main

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
)

type SQLiteKeyRepo struct {
	DB *sql.DB
}

func (kr *SQLiteKeyRepo) PrepDB() error {
	_, err := kr.DB.Exec(`CREATE TABLE IF NOT EXISTS keys(
		id INTEGER PRIMARY KEY,
		private_key TEXT,
		public_key TEXT,
		created_at TEXT
	)`)

	return err
}

func (kr *SQLiteKeyRepo) StoreKey(privateKey *rsa.PrivateKey) (string, error) {
	query := `INSERT INTO keys (
		private_key,
		public_key,
		created_at)
		VALUES (
			?,
			?,
			?)`

	serialPrivateKey := serializePrivateKey(privateKey)
	serialPublicKey, err := serializePublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", fmt.Errorf("error serializing public key: %+v", err)
	}
	createdAt := time.Now().Format(time.RFC3339)
	_, err = kr.DB.Exec(query, serialPrivateKey, serialPublicKey, createdAt)
	if err != nil {
		return "", fmt.Errorf("error storing new key: %+v", err)
	}
	return serialPublicKey, nil
}

func serializePrivateKey(key *rsa.PrivateKey) string {
	serialized := x509.MarshalPKCS1PrivateKey(key)
	encoded := base64.StdEncoding.EncodeToString(serialized)
	return encoded
}

func parsePrivateKey(encoded string) (*rsa.PrivateKey, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("error decoded base64 private key: %+v", err)
	}
	key, err := x509.ParsePKCS1PrivateKey(decoded)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %+v", err)
	}

	return key, nil
}

func serializePublicKey(key *rsa.PublicKey) (string, error) {
	serialized, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", fmt.Errorf("error marshalling public key: %+v", err)
	}
	encoded := base64.StdEncoding.EncodeToString(serialized)
	return encoded, nil
}

func parsePublicKey(encoded string) (*rsa.PublicKey, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("error decoded base64 private key: %+v", err)
	}
	key, err := x509.ParsePKIXPublicKey(decoded)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %+v", err)
	}

	ok := key.(*rsa.PublicKey)
	if ok == nil {
		return nil, errors.New("public key is not an rsa public key")
	}
	return ok, nil
}
