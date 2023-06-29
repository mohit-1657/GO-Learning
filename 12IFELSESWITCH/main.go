package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("now we learn if else conditions")

	loginCount := 12

	if loginCount < 10 {
		fmt.Println("regular user ")
	} else if loginCount > 10 {
		fmt.Println("supsious activity watch out")
	} else if loginCount < 0 {
		fmt.Println("No activity yet")
	} else {
		fmt.Println("kuch to kr ke bhai")
	}

	for i := 0; i < 10; i++ {
		if i&1 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	// another method where assigning the value and checking for that too at the same time

	if i := 3; i < 10 {
		fmt.Println("Number is less than the 10")
	} else {
		fmt.Println("numwer is graeter than the 10/")
	}

	// if err!=nil{
	// 	fmt.Println("",err)
	// }else{

	// }

	// lets start with the switch case

	fmt.Println("switch case  we will start learning")

	rand.Seed(time.Now().Unix())

	diceNumebr := rand.Intn(6) + 1

	fmt.Println("dice numner is ", diceNumebr)
	switch diceNumebr {
	case 1:
		{
			fmt.Println("you can open goti now")
		}
	case 2:
		{
			fmt.Println("can move to 2 place")
		}
	case 3:
		{
			fmt.Println("can move to 2 place")
		}
	case 4:
		{
			fmt.Println("can move to 2 place")
		}
	case 5:
		{
			fmt.Println("can move to 2 place")
		}
	case 6:
		{
			fmt.Println("can move to 6 place or open you goti ")
		}
	default:
		{
			fmt.Println("shi se chl na be")
		}
	}
}
