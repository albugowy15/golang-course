package main

import "fmt"

func main()  {
	defer fmt.Println("Hello World 1")
	defer fmt.Println("Hello World 2")
	defer fmt.Println("Hello World 3")
	fmt.Println("Learn defer")
	myDefer()

}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}