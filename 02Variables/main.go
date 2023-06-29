package main

import "fmt"

// in go we declare public varivale with the first character is capital;
var PublicVar int = 12

const ConstantVar string = "wttoken"

func main() {

	// in goif you are not  using variables which you have declared  then it will throw  an error for not doing good pratice

	var userName string = "rajeevtyagii80"

	fmt.Println(userName)
	fmt.Printf("This will tell us the type of  data : %T \n", userName)

	var isLoggedIn bool = false

	fmt.Printf("This will tell us the type of  data : %T \n", isLoggedIn)

	var smallint uint8 = 124
	fmt.Printf("This will tell us the type of  data : %T \n", smallint)

	var floatingPoint float32 = 26.84849948494844
	fmt.Println(floatingPoint)

	var floatingPoint64 float64 = 26.84849948494844444444
	fmt.Println(floatingPoint64)

	// variable declaration without specifying datatype
	var name = "Hello Boss"
	fmt.Println(name)

	var integerValue = 13
	fmt.Println(integerValue)

	// another way of declaration with data type specificationa and var

	withoutVarDeclaration := 12
	fmt.Println(withoutVarDeclaration)

	withoutVarDeclarationS := "without var declaration"

	fmt.Println(withoutVarDeclarationS)

	fmt.Println("Public variable is declared with first char is capital ", PublicVar)

	//ConstantVar ="fbfhb"   // constant variable can not be intilized again
	fmt.Println(ConstantVar)

}
