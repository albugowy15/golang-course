package main

import "fmt"

const LoginToken string ="6346346343"
func main()  {
	var username string  = "bughowi"
	fmt.Println("Hello from", username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool  = false
	fmt.Println("Hello from", isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8  = 255
	fmt.Println("Hello from", smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32  = 255.6776565565
	fmt.Println("Hello from", smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values
	var defautlValueVar int
	fmt.Println("Hello from", defautlValueVar)
	fmt.Printf("Variable is of type: %T \n", defautlValueVar)

	// Implicit type
	var implicitTypeVar = "Implicit type"
	fmt.Println("Hello from", implicitTypeVar)
	fmt.Printf("Variable is of type: %T \n", implicitTypeVar)

	// no var keyword
	noVarKeyword := "No var keyword"
	fmt.Println("Hello from", noVarKeyword)
	fmt.Printf("Variable is of type: %T \n", noVarKeyword)

	fmt.Println("Login Token is", LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	
}