package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dummyjson.com/products"
func main()  {
	fmt.Println("Learn Web Requests in Go")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type is %T\n", response)
	defer response.Body.Close() // close the connection

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}