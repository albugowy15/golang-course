package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://dummyjson.com/products?limit=10&skip=10&select=title,price"
func main()  {
	fmt.Println("Learn url in Go")
	fmt.Println(myurl)

	//parsing
	result, _ :=url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of qparams is %T\n", qparams)
	fmt.Println(qparams["skip"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "dummyjson.com",
		Path: "/products",
		RawQuery: "limit=10&skip=10&select=title,price",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}