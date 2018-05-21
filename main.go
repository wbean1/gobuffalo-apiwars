package main

import (
	"log"

	"github.com/wbean1/apiwars/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
