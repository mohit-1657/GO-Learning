package main

import "fmt"

func main() {
	fmt.Println("we will work on loops ")

	// slices

	days := []string{"Monday", "Tuesday", "wednesday", "thrusday", "Friday", "Statuday", "Sunday"}

	// iterarating with some sort of conditions which which can be amnipulated

	fmt.Println("Normal Loops itraration")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// for when we have to itertrate through the complete list then we can use
	// for each loop
	// in other languages while in case of for each it return the values but in go it returns index

	// fmt.Println("For each loop iteratation")
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for val string:days{
	// 	fmt.Println("",val)
	// }

	// same time with index and value return
	// fmt.Println("with both index and values")
	// for index, value := range days {
	// 	fmt.Println("", index, "    ", value)
	// }

	// with break and continue
	// rougeValue := 2

	// for rougeValue < 10 {
	// 	fmt.Println("value is now :", rougeValue)

	// 	// if rougeValue == 6 {
	// 	// 	break
	// 	// }

	// 	// if rougeValue == 6 {
	// 	// 	continue
	// 	// }  // tthis will took me again in infintie

	// 	// 	if rougeValue == 6 {
	// 	// 		rougeValue++
	// 	// 		continue
	// 	// 	}

	// 	// 	rougeValue++

	// 	// }
	// }
	// when we don't want to use the index only the values then

	for _, value := range days {

		if value == "wednesday" {
			goto lco
			continue
		}
		fmt.Println("values is in the list", value)
	}

	// now usng some labels

lco:
	fmt.Println("now you jumping at the labelled code")
}
