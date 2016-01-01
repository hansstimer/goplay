package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		fmt.Printf("rsa.GenerateKey: %v\n", err)
	}

	message := "Hello World!"
	messageBytes := bytes.NewBufferString(message)
	sha1 := sha1.New()

	encrypted, err := rsa.EncryptOAEP(sha1, rand.Reader, &privateKey.PublicKey, messageBytes.Bytes(), nil)
	if err != nil {
		fmt.Printf("EncryptOAEP: %s\n", err)
	}

	decrypted, err := rsa.DecryptOAEP(sha1, rand.Reader, privateKey, encrypted, nil)
	if err != nil {
		fmt.Printf("decrypt: %s\n", err)
	}

	decryptedString := bytes.NewBuffer(decrypted).String()
	fmt.Printf("message: %v\n", message)
	fmt.Printf("encrypted: %v\n", encrypted)
	fmt.Printf("decryptedString: %v\n", decryptedString)
}
