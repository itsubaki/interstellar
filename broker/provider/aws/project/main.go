package main

import (
	"log"
	"os"

	"github.com/itsubaki/interstellar/broker"
)

func main() {
	b, err := NewProjectBroker()
	if err != nil {
		log.Printf("new broker: %v", err)
		os.Exit(1)
	}

	broker.Run(b)
}
