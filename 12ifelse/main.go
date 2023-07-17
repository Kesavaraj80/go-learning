package main

import "fmt"

func main() {
	fmt.Println("IF ELSE in Go Lang")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "High User count"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is not less then 10")
	}

}
