package main

import (
	"fmt"
	"log"
	"reyhan/module"
)

func main() {
	messages, err := module.GenerateMultipleHelloMessage([]string{"John", "Jeremy", "Steve"})

	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		fmt.Println(name + ": " + message)
	}
}
