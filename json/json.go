package main

import (
	"encoding/json"
	"log"
)

type Foo struct {
	Count int
	Name  string
}

func main() {
	a := []Foo{{1, "pear"}, {2, "apple"}, {4, "banana"}}
	b := make([]Foo, 0)

	bytes, err := json.Marshal(a)
	if err != nil {
		log.Fatalf("Marshal failed:", err)
	}

	log.Printf("premarshal: %v\n", a)
	log.Println("json:", string(bytes))

	err = json.Unmarshal(bytes, &b)
	if err != nil {
		log.Fatalf("Unmarshal failed: %v\n", err)
	}
	log.Printf("unmarshalled: %v\n", b)
}
