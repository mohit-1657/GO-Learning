package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's start maps")

	language := make(map[string]string)

	language["JS"] = "Javascripts"
	language["RB"] = "RUBY"
	language["PY"] = "PYTHON"

	fmt.Println("printing the value for map :", language, len(language))

	delete(language, "RB")

	fmt.Println("printing the value for map :", language, len(language))

	// iteratating for maps

	fmt.Println("printing through the loops")
	for key, value := range language {
		fmt.Println(key, "   :", value)
	}

	fmt.Println("checking types of the maps variables and value")
	for key, value := range language {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

}
