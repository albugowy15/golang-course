package main

import "fmt"

func main() {
	fmt.Println("Learn structs")
	// no inheritance in go; no super class

	bughowi := User{"Bughowi", "bughowi@gmail.com", true, 22}
	fmt.Println(bughowi)
	fmt.Printf("bughowi details are: %+v\n", bughowi)
	fmt.Printf("Name is %v and email is %v\n", bughowi.Name, bughowi.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}