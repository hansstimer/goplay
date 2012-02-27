package main

import (
	"json"
	"log"
)

type Foo struct {
	count int
	name string
}

var a = []Foo{ {1, "pear"}, {2, "apple"}, {4, "banana"} }

func main() {
	b, err := json.Marshal(a)
	if err != nil {
		log.Exitln("Marshal failed:", err)
	}
	
	log.Printf("%v\n", a)
	log.Println("json:", string(b))
	
	
}

