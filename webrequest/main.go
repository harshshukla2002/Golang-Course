package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://harshshukla2002.github.io/"

func main() {
	fmt.Println("Welcome to webrequest in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Reponse type is %T\n", response)

	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data is: ", string(databyte))

}
