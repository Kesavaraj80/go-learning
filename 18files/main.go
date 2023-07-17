package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Go lang")
	content := "This needs to go in a file"

	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkError(err)

	length, err := io.WriteString(file, content)

	checkError(err)

	fmt.Println("Length is :", length)

	defer file.Close()
	readFile("./mygofile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkError(err)

	fmt.Println("Text data inside the file is\n", string(dataByte))
}

// Error handling function

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
