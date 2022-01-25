package main

import (
	"autobubble/api"
	"log"
)

func main() {
	_, err := api.GetEntityStates()
	if err != nil {
		log.Fatal(err)
	}
}
