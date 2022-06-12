package main

import (
	"log"
	"os"

	"github.com/depri11/vehicle/src/configs/command"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
