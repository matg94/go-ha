package main

import (
	"autobubble/api"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, world!")
	apiStatus, err := api.GetAPIState()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(apiStatus)
}
