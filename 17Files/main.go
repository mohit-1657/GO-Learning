package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// we check here for reading writting into files some multiple functanlities

	fmt.Println("Let's start with files")

	content := "I am writting this content into files for just having the fun"

	// file ,err =os.Create("directory path")
	file, err := os.Create("./myFirstFileCheck.txt")
	// file might be created or it might through some error
	if err != nil {
		panic(err) // panic is used to shut down the program oif error encountered
	}

	// now writting the content into files

	// so for thta we io package

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length of the files is : ", length)
	defer file.Close()
	readFiles("./myFirstFileCheck.txt")
}

func readFiles(pathwithFileName string) {

	textByte, err := ioutil.ReadFile(pathwithFileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("the text value int the file is : ", textByte) // it will give the data into the byte format

	fileContent := string(textByte)

	fmt.Println("readed content from the file is : ", fileContent)
}
