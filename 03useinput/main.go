package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	//comma ok//err ok
	input, _ := reader.ReadString('\n') //use "_" if you don't want to use the variable
	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("Type of input is: %T\n\n", input)

}
