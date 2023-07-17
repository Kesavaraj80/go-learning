package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Strawberry"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Fruit list length is", len(fruitList))

	var vegList = [3]string{"Tomato", "beans", "onions"}

	fmt.Println("Veg list", vegList)
}
