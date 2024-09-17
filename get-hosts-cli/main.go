package main

import (
	"get-hosts-cli/app"
	"log"
	"os"
)

func main() {
	application := app.Process()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
