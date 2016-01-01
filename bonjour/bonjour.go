package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oleksandr/bonjour"
)

func main() {
	resolver, err := bonjour.NewResolver(nil)
	if err != nil {
		log.Println("Failed to initialize resolver:", err.Error())
		os.Exit(1)
	}

	results := make(chan *bonjour.ServiceEntry)

	t1 := time.Now()
	go func(results chan *bonjour.ServiceEntry, exitCh chan<- bool) {
		for e := range results {
			log.Printf("%s", e.Instance)
			fmt.Printf("duration: %s\n", time.Now().Sub(t1)/time.Millisecond)
			exitCh <- true
			time.Sleep(1e9)
			os.Exit(0)
		}
	}(results, resolver.Exit)

	err = resolver.Browse("_ssh._tcp", "local.", results)
	if err != nil {
		log.Println("Failed to browse:", err.Error())
	}

	select {}
}
