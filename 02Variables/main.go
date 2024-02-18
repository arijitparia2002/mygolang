package main

import "fmt"

const LoginToken string = "hdjkfkajfhnfashdadhfd" //public "first letter capital"

func main() {
	fmt.Println("Variables")
	var username string = "Arijit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is of type: %T \n\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of smallval: %T\n\n", smallVal)

	var smallFloat float32 = 255.47348327432
	fmt.Println(smallFloat)
	fmt.Printf("smallFloat is of Type: %T\n\n", smallFloat)

	//defult values and some aliases
	var anothervalue int
	fmt.Println(anothervalue)
	fmt.Printf("Type of anothervarible is: %T\n\n", anothervalue)

	var website = "TUTORIEX"
	fmt.Println(website)

	//no var style
	numberofuser := 4000
	fmt.Println(numberofuser)

	fmt.Println(LoginToken)

}
