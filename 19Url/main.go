package main

import (
	"fmt"
	"net/url"
)

func main() {

	urls := "htpps://lco.dev:3000/learn?coursename=reactjs&payment=jfnfnhefbhfbhrfb"
	handelingUrls(urls)
}

func handelingUrls(myUrl string) {

	fmt.Println("My urls :", myUrl)
	fmt.Println("now going throuyght the different url configurations")

	// paring the url

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}
	// fmt.Println("scheme :", result.Scheme)
	// fmt.Println("Host   :", result.Host)
	// fmt.Println("path   :", result.Path)
	// fmt.Println("Post   :", result.Port()) // post is basically methods
	fmt.Println("Query  :", result.RawQuery)

	// now we can get complete query param

	queryParam := result.Query()

	fmt.Printf("type the query param is : %T\n", queryParam)

	// in that  we can see that the queryparam  store the value in key value in a map

	fmt.Println("all param ", queryParam)

	fmt.Println("the valye for key [coursename] is ", queryParam["coursename"])

	for key, value := range queryParam {
		fmt.Println("for key ", key, " value is ", value)
	}

	// craeting the url in different way

	partsOfurl := url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	// another url construct
	anotherUrl := partsOfurl.String()

	fmt.Println("", anotherUrl)
}
