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

	hitesh.getuserStatus()
	fmt.Println()
}

// we putted first character capital because it will access publicly from every where

// struts are same behaved as the classed
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int

	// if made the varibales with first car is small then it will be exportable in different class
	//oneAge int
}

func (u User) getuserStatus() bool {

	fmt.Println("hello the status of user is ", u.Status)
	return u.Status

}

// setting the value for struct out side
// but setting the value like this thois will only save till the fuction because our  manipulation
// is on the copy of object  when we set on the reference of variale then it will for the all
func (u User) setuserStatus() bool {

	u.Status = false
	fmt.Println("user status is now ", u.Status)
	return u.Status

}
