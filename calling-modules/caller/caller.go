package main

import (
	"fmt"
	"reyhan/module"
)

func main() {
	message := module.GenerateHelloMessage("Reyhan")
	fmt.Println(message)
}
