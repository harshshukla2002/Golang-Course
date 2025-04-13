package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in golang")

	// there is not inheritance in golang: no super or parent.
	harsh := User{"Harsh", "harsh@mail.com", true, 21}
	harsh.GetStatus()
	harsh.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("email of this user is:  ", u.Email)
}
