package main

import "fmt"

func main() {
	fmt.Println("Welconme to array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Bananan"
	fruitList[3] = "Orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	//veg list
	var vegList = [3]string{"potato", "mashroom", "beans"}

	fmt.Println("Data -->", vegList, "Length is:", len(vegList))

}
