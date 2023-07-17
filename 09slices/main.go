package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Orange", "Peach"}
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 111
	highScores[1] = 222
	highScores[2] = 333
	highScores[3] = 444
	// this will not work
	// highScores[4]=555

	// append will work
	highScores = append(highScores, 555, 666, 777, 888)

	fmt.Println(highScores)

	// how to remove value from slice based on index
	var courses = []string{"React", "Java", "Python", "swift", "Ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
