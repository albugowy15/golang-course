package main

import "fmt"

func main()  {
	fmt.Println("Learn maps")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"

	fmt.Println(languages)
	fmt.Println("GO ", languages["GO"])

	delete(languages, "RB")
	fmt.Println(languages)

	// looping in maps
	for key, value := range languages {
		fmt.Println(key, value)
	}
}