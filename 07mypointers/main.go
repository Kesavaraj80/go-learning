package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 25

	var ptr = &myNumber
	fmt.Println("Address of myNumber using &", ptr)
	fmt.Println("value of myNumber by accesing via pointer using *:", *ptr)

	*ptr = *ptr + 10

	fmt.Println("New value of myNumber", myNumber)
}
