package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	//get the current time
	now := time.Now()
	fmt.Println(now)

	//format the time
	formattedTime := now.Format("02-01-2006 15:04 Monday") //remember this
	println(formattedTime)

	//create date
	createDate := time.Date(2020, time.August, 13, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
