package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Tomato"} //slices

	fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Banana", "Mango", "Orange")
	fmt.Println("Updated fruitList:", fruitList)

	// Here you need to provide the values you want to append
	// For example, if you want to append "Pineapple" and "Grape", you can do:
	fruitList = append(fruitList[1:], "Pineapple", "Grape")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:5])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 90
	highScore[1] = 390
	highScore[2] = 323
	highScore[3] = 23232
	// highScore[4] = 34 //panic ==> error

	highScore = append(highScore, 211, 900) //

	fmt.Println(highScore)

	sort.Ints(highScore) //sorting integer
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

}
