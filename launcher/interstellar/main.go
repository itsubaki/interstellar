package main

import (
	"log"

	"github.com/itsubaki/interstellar/launcher"
)

func main() {
	i := NewInterstellar()
	if err := launcher.Run(i); err != nil {
		log.Println(err)
	}
}
