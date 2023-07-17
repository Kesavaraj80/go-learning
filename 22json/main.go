package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Creating JSON in GO")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 100, "LCO", "test@123", []string{"web-dev", "js"}},
		{"NodeJS", 79, "LCO", "test@123", []string{"backend-dev", "js"}},
		{"NextJS", 200, "LCO", "test@123", []string{"frontend-dev", "js"}},
	}

	// Package this data as JSON Data

	finaljson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)

}

func DecodeJson() {
	jsonDatefromWeb := []byte(`
	{
		"coursename": "ReactJS",
		"Price": 100,
		"Platform": "LCO",
		"tags": ["web-dev","js"]
	}
	`)

	var lcocourse course

	checkValid := json.Valid(jsonDatefromWeb)

	if checkValid {
		fmt.Println("JSON Was valid")
		json.Unmarshal(jsonDatefromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("Json not valid")
	}

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDatefromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is:%T\n", k, v, v)
	}
}
