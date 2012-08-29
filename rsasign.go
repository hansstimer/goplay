package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("rsa.GenerateKey: %v\n", err)
		return
	}

	message := "Hello World!"
	messageBytes := bytes.NewBufferString(message)
	hash := sha512.New()
	hash.Write(messageBytes.Bytes())
	digest := hash.Sum(nil)

	fmt.Printf("messageBytes: %v\n", messageBytes)
	fmt.Printf("hash: %V\n", hash)
	fmt.Printf("digest: %v\n", digest)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, digest)
	if err != nil {
		fmt.Printf("rsa.SignPKCS1v15 error: %v\n", err)
		return
	}

	fmt.Printf("signature: %v\n", signature)

	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA512, digest, signature)
	if err != nil {
		fmt.Printf("rsa.VerifyPKCS1v15 error: %V\n", err)
	}

	fmt.Println("Signature good!")

}
