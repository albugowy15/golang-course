package main

import "fmt"

func main()  {
	fmt.Println("Hello World")
	var fruitList [4]string
	
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	// fruitList[2] = "Grapes"
	fruitList[3] = "Mango"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("Veg list is : ", vegList)
}