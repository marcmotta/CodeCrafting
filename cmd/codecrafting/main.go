// cmd/codecrafting/main.go
package main

import (
	"flag"
	"log"
	"os"

	"codecrafting/internal/codecrafting"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := codecrafting.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
