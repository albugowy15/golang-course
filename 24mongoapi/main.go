package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albugowy15/mongoapi/router"
)

func main()  {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is gettin started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}