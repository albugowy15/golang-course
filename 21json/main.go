package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main()  {
	fmt.Println("Learn json")
	DecodeJson()
}

func EncodedJson()  {
	mycourses := []course{
		{"React Bootcamp",299, "bughowi.com", "abc123", []string{"web-dev", "js"}},
		{"Angular bootcamp",299, "bughowi.com", "774hh4", []string{"fullstack-dev", "ts"}},
		{"Vue Bootcamp",299, "bughowi.com", "hdhd72", nil},
	
	}

	// package this data as JSON data

	finalJson, err :=json.MarshalIndent(mycourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "React Bootcamp",
			"Price": 299,
			"website": "bughowi.com",
			"tags": ["web-dev","js"]
		}
	`)

	var myCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Valid Json")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("Invalid Json")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is %T\n",k, v, v)
	}
	
}