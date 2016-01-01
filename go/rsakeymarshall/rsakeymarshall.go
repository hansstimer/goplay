package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"log"
)

type MyPrivateKey struct {
	privateKey *rsa.PrivateKey
}

func (p MyPrivateKey) MarshalJSON() ([]byte, error) {
	priv, err := json.Marshal(p.privateKey)
	if err != nil {
		log.Fatalf("json.Marshal: %v\n", err)
	}

	pub, err := json.Marshal(p.privateKey.PublicKey)
	if err != nil {
		log.Fatalf("json.Marshal: %v\n", err)
	}

	bytes := append([]byte(`{ "PrivateKey": `), priv...)
	bytes = append(bytes, []byte(`,"PublicKey": `)...)
	bytes = append(bytes, pub...)
	bytes = append(bytes, []byte(" }")...)

	return bytes, nil
}

func main() {
	var privateKey MyPrivateKey
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v\n", err)
	}
	privateKey.privateKey = pk

	bytes, err := json.Marshal(privateKey)
	if err != nil {
		log.Fatalf("Marshal failed: %v\n", err)
	}

	log.Printf("privateKey: %V\n", privateKey)
	log.Println("json privateKey:", string(bytes))
}
