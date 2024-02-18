package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointer")

	// var ptr *int
	// fmt.Println("Value of ptr is:", ptr) //<nil>

	myNumber := 32
	var ptr *int = &myNumber
	fmt.Println("Value of the actual ptr is:", ptr)
	fmt.Println("Value inside of the actual ptr is:", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumber)

}
