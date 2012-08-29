package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"hash"
)

func encrypt(hash hash.Hash, publicKey *rsa.PublicKey, in []byte) ([]byte, error) {
	out, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, in, nil)
	if err != nil {
		return nil, err
	}

	return out, err
}

func decrypt(hash hash.Hash, privateKey *rsa.PrivateKey, in []byte) ([]byte, error) {
	out, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, in, nil)
	if err != nil {
		return nil, err
	}

	return out, err
}

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		fmt.Printf("rsa.GenerateKey: %v\n", err)
	}

	message := "Hello World!"
	messageBytes := bytes.NewBufferString(message)
	sha1 := sha1.New()

	encrypted, err := encrypt(sha1, &privateKey.PublicKey, messageBytes.Bytes())
	if err != nil {
		fmt.Printf("encrypt: %s\n", err)
	}

	decrypted, err := decrypt(sha1, privateKey, encrypted)
	if err != nil {
		fmt.Printf("decrypt: %s\n", err)
	}

	decryptedString := bytes.NewBuffer(decrypted).String()
	fmt.Printf("message: %v\n", message)
	fmt.Printf("encrypted: %v\n", encrypted)
	fmt.Printf("decryptedString: %v\n", decryptedString)
}
