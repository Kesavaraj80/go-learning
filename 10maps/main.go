package main

import "fmt"

func main() {
	fmt.Println("Maps in go")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	// loops

	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}
}
