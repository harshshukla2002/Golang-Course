package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["CPP"] = "CPlusPlus"
	languages["JSX"] = "JavaScript Xml"

	fmt.Println("Map", languages)

	// delete
	delete(languages, "JSX")
	fmt.Println("List of all languages after deletion", languages)
}
