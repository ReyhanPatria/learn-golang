package main

import (
	"fmt"
	"log"
	"reyhan/module"
)

func main() {
	message, err := module.GenerateHelloMessage("Reyhan")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
