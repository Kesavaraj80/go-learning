package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in go Lang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is", result)
	proResult, message := proAdder(2, 3, 4, 1, 6, 7, 8)

	fmt.Println("Value is", proResult)
	fmt.Println("Message is", message)
	// greeteTwo()
}

// func greeteTwo() {
// 	fmt.Println("Another function")
// }

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Returning two values"
}

func greeter() {
	fmt.Println("Vankkam from go lang")
}
