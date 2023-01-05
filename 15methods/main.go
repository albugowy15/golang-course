package main

import "fmt"

func main() {
	fmt.Println("Learn methods")
	// no inheritance in go; no super class

	bughowi := User{"Bughowi", "bughowi@gmail.com", true, 22}
	fmt.Println(bughowi)
	fmt.Printf("bughowi details are: %+v\n", bughowi)
	fmt.Printf("Name is %v and email is %v\n", bughowi.Name, bughowi.Email)
	bughowi.GetStatus()
	bughowi.NewMail()
	fmt.Printf("Name is %v and email is %v\n", bughowi.Name, bughowi.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int

}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is", u.Email)
} 