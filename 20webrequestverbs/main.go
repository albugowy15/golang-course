package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	fmt.Println("Learn get request golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
	
}

func PerformGetRequest() {
	const myUrl = "http://localhost:1111/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code is", response.StatusCode)
	fmt.Println("Content length is", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is", byteCount)
	fmt.Println("Response string is", responseString.String())
	// fmt.Println(content)
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:1111/post"

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename": "Lets learn golang",
		"price": 0,
		"platform": "Youtube"
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content)) 
}

func PerformPostFormRequest()  {
	const myUrl = "http://localhost:1111/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Mohamad")
	data.Add("lastname", "bughowi")
	data.Add("email", "kholidbughowi@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}