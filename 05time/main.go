package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to time study in golang")

	presentTime := time.Now()
	fmt.Println("Present time is ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.October, 10, 22, 20, 0, 0, time.UTC)
	fmt.Println("Created date is ", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}