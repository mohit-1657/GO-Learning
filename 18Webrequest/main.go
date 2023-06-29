package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	myfristWebRequest("https://lco.dev")
}

func myfristWebRequest(url string) {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // this is caller reponsibility to close

	// now checking type of response

	fmt.Printf("type the web request is %T\n", response)
	// in re type of reponse you will see a pointer (*http.Response) which showing it is returningthe original reference not copy

	// reading the content
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println("the content for the response body is :", content)

}
