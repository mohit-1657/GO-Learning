package main

import (
	"fmt"
)

// there is no inheritance in the go, nio concepts of super keywords or parent concepts

func main() {
	fmt.Println("Starts with structs in go lang")

	hitesh := User{"Hitesh", "hitesh.com", true, 34}

	fmt.Println("hitesh details :", hitesh)
	fmt.Println("printing the types for the structs")
	fmt.Printf("Name type %+v\n", hitesh.Name)
	fmt.Printf("Email type: %v status type : %v Age type %v\n", hitesh.Name, hitesh.Status, hitesh.Age)
}

// we putted first character capital because it will access publicly from every where
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
