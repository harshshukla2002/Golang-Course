package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println("Time is:", presentTime.Format("01/02/2006 03:04:05PM Monday"))
}
