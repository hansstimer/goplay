package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
)

func encrypt(privateKey *rsa.PrivateKey, in []byte) {

}

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		fmt.Printf("rsa.GenerateKey: %v", err)
	}
	sha1 := sha1.New()

	message := "Hello World!"
	in := bytes.NewBufferString(message)
	out, err := rsa.EncryptOAEP(sha1, rand.Reader, &privateKey.PublicKey, in.Bytes(), nil)
	if err != nil {
		fmt.Printf("rsa.EncryptOAEP: %s", err)
	}

	fmt.Printf("%V", out)
}
