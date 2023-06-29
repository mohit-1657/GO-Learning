package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome for the function class")
	// greeting("Afternoon")

	// greeting1("Afternoon")
	// greeting3("Afternoon")

	result := adder(23, 55)

	fmt.Println("result of two value added :", result)

	// result = proAdder(10, 20, 20)
	// fmt.Println("result with number of variables", result)
	// fmt.Println("result with number of variables", proAdder(20, 30, 40, 50))

	value1, value2 := differentDataReturn(20, "Ram, ")

	fmt.Println(value2, value1)
}

func greeting(greetingTime string) {

	fmt.Println("Good ", greetingTime)
}

// inside any function we can not define any new function

func greeting1(greetingTime string) {

	// func greeting2(greetingTime string) {

	// 	fmt.Println("Good ", greetingTime)
	// }
	fmt.Println("as Good as hell", greetingTime)
}

func greeting3(greetingTime string) {

	fmt.Println("calling from function 3 ")
	greeting1(greetingTime + "3")

	fmt.Println("Good ", greetingTime)
}

// annoymous function can also be declared

// func(greetingTime string){

// 	fmt.Println("Hr",)
// }

// method witkh multiples variables and getting return in different type

// decalaration of fuction when we have a particular return type
// func signature
// func name(,,,,) returntype{

// }
func adder(num1 int, num2 int) int {

	return num1 + num2
}

//function signature when we don't know how many variables are comings

func proAdder(values ...int) int {

	total := 0

	// we can that the values are came in the form of slices or we can say that it is a type of array or array list
	// so to get the values we have to iteratate or getting on the idex bases
	for _, value := range values {
		total += value
	}

	return total

	// we can return 2 diffterent datatype from single function

}

func differentDataReturn(val1 int, val2 string) (int, string) {

	return val1 + 10, val2 + "hello you are getting values from the fucntion"
}
