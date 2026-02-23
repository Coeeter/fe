package main

import (
	"log"

	"github.com/coeeter/fe/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal("Error executing command:", err)
	}
}
