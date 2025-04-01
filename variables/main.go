package main

import "fmt"

func main() {
	var username string = "Harsh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue int = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float32 = 12.44454555454545545
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 12.44454555454545545
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	var undefinedInt int
	fmt.Println(undefinedInt)
	fmt.Printf("Variable is of type: %T \n", undefinedInt)

	var undefinedBool bool
	fmt.Println(undefinedBool)
	fmt.Printf("Variable is of type: %T \n", undefinedBool)

	var undefinedString string
	fmt.Println(undefinedString)
	fmt.Printf("Variable is of type: %T \n", undefinedString)

	var undefinedFloat float32
	fmt.Println(undefinedFloat)
	fmt.Printf("Variable is of type: %T \n", undefinedFloat)

	var name = "Harsh"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	numberOfUser := 23
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)
}
