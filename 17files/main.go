package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Learn files")

	content := "This needs to go in a file"

	file, err := os.Create("./myfile.txt")
	checkNiErr(err)

	length, err := io.WriteString(file, content)
	checkNiErr(err)
	fmt.Println("Length of file is", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNiErr(err)
	fmt.Println("Text data inside the file is", string(databyte))
}

func checkNiErr(err error) {
	if err != nil {
		panic(err)
	}
}