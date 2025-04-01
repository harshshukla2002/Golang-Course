package main

import "fmt"

func main() {
	fmt.Println("---------Pointer--------")

	// var one *int
	// fmt.Println("Value of pointer is", one)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("location of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value of mynumber is", myNumber)
}
