package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome on Files Operations in Golang")

	content := "This is created by files operations"

	file, err := os.Create("./newText.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("File is created and Length is", length)

	defer file.Close()

	readFile("./newText.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside file is: ", string(databyte))
}
