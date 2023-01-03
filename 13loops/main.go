package main

import "fmt"

func main()  {
	fmt.Println("Learn loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d:=0; d<len(days); d++ {
	// 	fmt.Println(days[d])
	
	// }

	// for i:= range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println(index, day)
	// }

	rougueValue := 1 

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5	 {
			rougueValue++
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

	lco:
		fmt.Println("Jumping at bughowi.com")
}