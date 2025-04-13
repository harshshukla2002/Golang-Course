package main

import "fmt"

func main() {
	fmt.Println("Learning Defer")

	defer fmt.Println("Hello")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("World")
}
