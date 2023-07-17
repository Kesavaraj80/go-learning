package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "User input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating:")

	// comma ok // err ok

	input, _ := reader.ReadString('\n')

	fmt.Println("your rating", input)
	fmt.Printf("type of rating is %T:", input)

}
