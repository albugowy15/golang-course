package main

import "fmt"

func main()  {
	fmt.Println("Lets learn pointer in go")
	// var ptr *int

	// fmt.Println("The value of ptr is ", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("The value of ptr is ", ptr)
	fmt.Println("The value of ptr is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New values is : ", myNumber)
}