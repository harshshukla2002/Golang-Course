package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in golang")
	// slice loop
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for d := 0; d < len(days); d++ {
		fmt.Println("days", days[d])
	}

	// another way
	for i := range days {
		fmt.Println(days[i])
	}

	// another way
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// another way
	index := 1

	for index < 5 {
		if index == 2 {
			goto lco
		}

		fmt.Printf("index is %v\n", index)
		index++
	}

	// goto keyword
lco:
	fmt.Println("This is goto keyword uses")

	// map loop
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["CPP"] = "CPlusPlus"
	languages["JSX"] = "JavaScript Xml"

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
}
