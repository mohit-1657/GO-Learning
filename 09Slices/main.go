package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	fmt.Println("let's start learning slices")
	var fruitList [4]string
	fruitList[0] = "werfg:"
	fruitList[3] = ":fnfnnfnfnf"
	fmt.Println(fruitList)

	var vegitableList = [3]string{"Hello", "Hello1"}
	fmt.Println(vegitableList)

	// above is array part

	// now slices

	var grocerry = []string{}

	//int this we can see we have intilized but not provided the fixed sizes
	//  this is same as the arraylist where we will save apped the data

	grocerry = append(grocerry, "potato", "tomato")
	fmt.Println(grocerry)

	// another way is that we can provide some values during the intilization
	var grocerry1 = []string{"apple", "mango"}
	fmt.Println("while intilization list :", grocerry1)
	// now if want addon then we have to use the append method

	grocerry1 = append(grocerry1, "carrot", "kiwi", "peach")
	fmt.Println("After appending the values in list :", grocerry1)

	// types check
	fmt.Printf("type is %T\n", grocerry1)

	// now slicling

	grocerry1 = append(grocerry1[1:]) // passing the index

	fmt.Println(grocerry1)
	grocerry1 = append(grocerry1[1:3])
	fmt.Println(grocerry1)

	// now defining the slices in different way

	// with make keyword

	highscore := make([]int, 4)

	highscore[0] = 222
	highscore[1] = 223
	highscore[2] = 224
	highscore[3] = 225

	// when you do this it gives you index out range error
	//highscore[0]=222

	// but with the append method it'll works

	//when we created slice with the make , intially it will defines the slices with only the given sizes
	// but when use append method it will reallocate the memory and initlized again the next append

	fmt.Println(highscore)

	highscore = append(highscore, 123, 388, 844)

	fmt.Println("after append method :", highscore)

	// one more thing to keep in mind is that the, at dafualt allocated memory we have to place the values  explicatillay by index based

	// append method will only at the end of the last the memory size

	// some others methods

	sort.Ints(highscore)

	//after the sorting
	fmt.Println("soreted ans  :", highscore)

	fmt.Println("is higscore list sorted ? :", sort.IntsAreSorted(highscore))

	for i := 0; i < len(highscore); i++ {
		fmt.Println(i, " : ", highscore[i])
	}

	writer := bufio.NewReader(os.Stdin)

	fmt.Println("Please Enter the marks")
	var input, _ = writer.ReadString('\n')

	fmt.Println("", input)

	// removing the ith index from the slices

	var index int = 1

	fmt.Println("highscore list before removing index", highscore)
	highscore = append(highscore[:index], highscore[index+1:]...)
	fmt.Println("highscore list after removing index", highscore)
	//.... is used because there are multiple values are passed in the method

	var studentsList = [2]string{}

	var studentList1 [3]string
	studentList1[1] = "fgvbhfbhbdfh"
	var studentList2 [3][3]int

	studentList2[0][1] = 1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			studentList2[i][j] = i * j
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(studentList2[i][j], " ")
		}
		fmt.Println("")
	}
	studentsList[0] = "dcbhvbhfd"

}
