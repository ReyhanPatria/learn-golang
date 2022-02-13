package main

import "reyhan/restful-api/model"

var products = []model.Product{
	{ID: 1, Name: "Apple", Price: 23000, Description: "Apple fresh from Aomori Prefecture"},
	{ID: 2, Name: "Banana", Price: 25000, Description: "Banana from the Banana Republic"},
	{ID: 3, Name: "Steam", Price: 0, Description: "It's the software, not actual steam"},
}

func main() {
}
