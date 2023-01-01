package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Learn slices")
	var fruitList = []string{"Apple", "Orange", "Grapes", "Mango"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fruitList = append(fruitList, "Pineapple", "Banana")
	fmt.Println("Fruit List is: ", fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println("Fruit List is: ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 234
	highScores[2] = 345
	highScores[3] = 456
	// highScores[4] = 8357

	// This will trigger a new memory allocation,
	// and copy the contents of the old slice to the new slice
	highScores = append(highScores, 34563, 38745, 5353)

	fmt.Println("High Scores are: ", highScores)

	sort.Ints(highScores)
	fmt.Println("High Scores are: ", highScores)
	fmt.Println("Is slices sorted", sort.IntsAreSorted(highScores))


	// remove a value from a slice wiht index
	var courses = []string{"reactjs", "typescript", "angularjs", "nodejs", "golang"}
	fmt.Println(courses)

	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}