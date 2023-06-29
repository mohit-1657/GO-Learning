package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello")

	timeParser := time.Now()

	fmt.Println(timeParser)

	fmt.Println(timeParser.Format("01-02-2006"))
	fmt.Println(timeParser.Format("2006-02-01"))
	fmt.Println(timeParser.Format("2006-01-02"))
	//timeParser(timeParser)

	// For the time formation in this is dd-mm-yyyy only "01-02-2006" this date will be passed
	//except that none of the other will not be accepted and that could be checked in the  package documentation

	// other date format coukd be check there too and each have a standard format input

	fmt.Println(timeParser.Format("2006-01-02 15:04:05 Monday"))

	presentDate := time.Date(2023, time.March, 6, 4, 34, 3, 3545, time.UTC)

	fmt.Println(presentDate)

	// these also could be formatted
	fmt.Println(presentDate.Format("01-02-2006 Monday"))
}
