package main

import (
	"fmt"
	"log"
	"reyhan/module"
)

func main() {
	message, err := module.GenerateHelloMessage(" ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
