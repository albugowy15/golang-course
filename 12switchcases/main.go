package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Learn switch cases")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice number is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got one and you can open")
	case 2:
		fmt.Println("You got two move spot")
	case 3:
		fmt.Println("You got three move spot")
	case 4:
		fmt.Println("You got four move spot")
		fallthrough
	case 5:
		fmt.Println("You got five move spot")
	case 6:
		fmt.Println("You got six move spot and roll again")
	default:
		fmt.Println("Invalid dice number")
	}
}