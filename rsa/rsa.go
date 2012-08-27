package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
)

func encrypt(hash hash.Hash, privateKey *rsa.PrivateKey, in []byte) ([]byte, error) {
	out, err := rsa.EncryptOAEP(hash, rand.Reader, &privateKey.PublicKey, in.Bytes(), nil)
	if err != nil {
		return nil, err
	}

	return out, err
}

func decrypt(hash hash.Hash, publicKey *rsa.PublicKey, in []byte) ([]byte, error) {
	out, err := rsa.DecryptOAEP(hash, rand.Reader, publicKey, in, nil)
	if err != nil {
		return nil, err
	}

	return out, err
}

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		fmt.Printf("rsa.GenerateKey: %v", err)
	}

	message := "Hello World!"
	messageBytes := bytes.NewBufferString(message)
	sha1 := sha1.New()

	out, err := encrypt(sha1, privateKey, messageBytes)
	if err != nil {
		fmt.Printf("encrypt: %s", err)
	}

	out, err := decrypt(sha1, &privateKey.PublicKey, out)
	if err != nil {
		fmt.Printf("encrypt: %s", err)
	}

	fmt.Printf("%V", out)
}
