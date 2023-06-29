package main

import "fmt"

func main() {
	fmt.Println("Let's start Arrays")

	var students [4]string

	students[0] = "Hello"
	students[1] = "Boy"
	students[3] = "Say yes"

	fmt.Println("what to say is : :", students)

	fmt.Println("length of array is :", len(students))
	// we will have a sigle space for the string type array
	// array creation and initiliztion at the same time

	var vegitableList = [3]string{"Potato", "tomamto", "onion"}

	fmt.Println("array creation and initiliztion at the same time:::::::", vegitableList)
}
