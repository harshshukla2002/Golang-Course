package main

import "fmt"

func main() {
	fmt.Println("Welcome in Arrays in Go Lang")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[3] = "Mango"

	fmt.Println("Fruits are", fruits)

	var veggies = [3]string{"Potato", "Tomato", "Onion"}

	fmt.Println("Veggeis are", veggies)
}
