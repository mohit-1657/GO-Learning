package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers")
	// when we pass on reference of the varibles in that case somtimes it passed the copy of thta variable and sometime irgegularities are occured,
	// to avoid those we use the pointers

	// it is direct reference of memory address

	var pointersRef1 *int
	fmt.Println("default value of pointers ", pointersRef1)

	// another way fpr the pointers

	value1 := 12

	pointersRef2 := &value1

	fmt.Println("the reference of the pointer value is annotted by the & ", pointersRef2)
	fmt.Println("the actual value of pointer is reffred by any pointer is ", *pointersRef2)
	//fmt.Println(pointersRef2)

	*pointersRef1 = 2 + *pointersRef1
	fmt.Println("the reference of the pointer value is annotted by the & ", pointersRef1)
	*pointersRef2 = *pointersRef2 * 2
	fmt.Println("the reference of the pointer value is annotted by the & ", pointersRef2)
	// if one the poiters is used in the variables then all variables will be with pointers else error
	*pointersRef2 = *pointersRef2 + 2
	fmt.Println("the reference of the pointer value is annotted by the & ", pointersRef2)

}
