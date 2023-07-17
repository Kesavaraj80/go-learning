package main

import "fmt"

func main() {
	fmt.Println("Loops in GO")

	// days := []string{"sunday", "Monday", "Tuesday", "wednesday", "Thursday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("index is %v and the value is %v\n", i, day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			break
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumped here")

}
