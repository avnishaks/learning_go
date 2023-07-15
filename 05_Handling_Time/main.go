package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study for the Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023,time.July,16,8,2,50,12,time.UTC)
	fmt.Println(createdDate)

}
