package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welocme the Go lang CLasses")
	fmt.Println("Pls give rating how much you hate learning a new language")

	reader := bufio.NewReader(os.Stdin)

	hateRating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your valuable rating ")
	fmt.Println(" consloing the hate rating ", hateRating)
	fmt.Printf("type of the given rating %T \n", hateRating)

	// data type conversion

	//numConverted, err := strconv.ParseFloat(hateRating, 64)
	// above line is going to throw an error beacause the /n will be added in the string as trailing character

	// to get rid of that
	numConverted, err := strconv.ParseFloat(strings.TrimSpace(hateRating), 64) // it will trim the space character from the laft and right

	fmt.Println("consoling converted variables ", numConverted)
	fmt.Printf("type of the given rating %T \n", numConverted)
	fmt.Println("Error while conerting ", err)

	if err != nil {

	} else {
		fmt.Println("Printing ")
	}
}
