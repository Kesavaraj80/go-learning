package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:06 Monday"))

	createdDate := time.Date(2020, time.April, 22, 10, 00, 00, 00, time.UTC)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 15:04:06 Monday"))
}
