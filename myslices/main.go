package main

import "fmt"

func main() {
	fmt.Println("Welcome to Slices in golang")

	var fruits = []string{"Mango", "Apple"} // initialise

	fruits = append(fruits, "Orange", "Banana") // adding value

	fruits = append(fruits[1:]) // getting value from given index

	fmt.Printf("Type of friuts is %T\n", fruits)
	fmt.Println("and length is", len(fruits))

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 454
	highScore[2] = 864
	highScore[3] = 914
	// highScore[4] = 1003

	highScore = append(highScore, 1003)

	fmt.Println("High scores are", highScore)

	// remove value form slice
	var courses = []string{"React", "Angular", "Java", "JavaScript", "Kotlin"}
	fmt.Println("Courses", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("after remove Courses", courses)
}
