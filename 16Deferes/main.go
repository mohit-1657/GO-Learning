package main

import "fmt"

func main() {
	defer fmt.Println("world")
	//defer fmt.Println("world1")
	fmt.Println("Hello")
	defer fmt.Println("world2")
	fmt.Println("hello1")

	mydefer()

}
func mydefer() {
	for i := 0; i < 5; i++ {
		//fmt.Print(i, ",")
		defer fmt.Print(i, ",")
	}
}
