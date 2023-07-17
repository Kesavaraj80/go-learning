package main

import "fmt"

func main() {
	fmt.Println("Welcome To structs in go lang")
	// no inheritence
	// no Super or Parent

	kesavan := User{"Kesavan", "Kesavan@go.dev", true, 23}

	fmt.Println(kesavan)
	fmt.Printf("Kesavan details are :%+v\n", kesavan)
	fmt.Printf("Name is %v and email is %v", kesavan.Name, kesavan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
