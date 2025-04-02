package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in golang")

	// there is not inheritance in golang: no super or parent.
	harsh := User{"Harsh", "harsh@mail.com", true, 21}
	fmt.Println("User is", harsh)
	fmt.Printf("Harsh details are: %+v\n", harsh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
