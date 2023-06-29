package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user inupt classs")

	reader := bufio.NewReader(os.Stdin) // in these methods we can some have starting with Caps Letter because thsese are public
	// coma ok  || error ok
	var rating, _ = reader.ReadString('\n')
	// value,_ (_) if any error  occured then it will be stored in the _
	fmt.Println("printing variable with the only storing input with non used error var", rating)

	// if you want to log the error then you can store that in a varibles

	var rating1, err = reader.ReadString('\n')

	fmt.Println("print values with the error variable ", rating1)
	fmt.Println("printing error ", err)
}
