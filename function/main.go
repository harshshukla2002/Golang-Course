package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in golang")
	greeter()
	fmt.Println(add(2, 3))
	fmt.Println(proAdder(1, 2, 3, 4, 5, 6))
}

func greeter() {
	fmt.Println("Heyy from greeter")
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}
