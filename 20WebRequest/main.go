package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	//performGETREquest()
	//performPOSTRequest()

	performPOSTformrequest()
}

func performGETREquest() {

	const myurls = "http://localhost:8000/get"

	reponse, err := http.Get(myurls)

	defer reponse.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("my request length is : ", reponse.ContentLength)
	fmt.Println("status code of request : ", reponse.StatusCode)
	fmt.Printf("fresponse type for the request %T\n", reponse)

	content, _ := ioutil.ReadAll(reponse.Body)
	fmt.Println("response  conttent ", content)

	//
	fmt.Println("response  conttent ", string(content))

	// using builder we can also used to write
	var responseString strings.Builder

	byteCount, _ := responseString.Write(content) // this will same as the count return
	// we have to change this againg in from the byte format

	fmt.Println("byte count is ", byteCount)
	newResponse := responseString.String()

	fmt.Println("response gerenerated with the builder :", newResponse)
}

func performPOSTRequest() {

	const postUrl = "http://localhost:8000/post"

	// making a json rrequest body

	requestBody := strings.NewReader(`
	{
		"courseName":"golang",
		"duration":"2 months",
		"price":111111
	}
	`)

	// making post request

	response, err := http.Post(postUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	//fmt.Println("", response)

	byteContent, _ := ioutil.ReadAll(response.Body)

	fmt.Println("", string(byteContent))

}

func performPOSTformrequest() {

	const postFornUrl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("Name", "Mohit")
	data.Add("latsName", "Kumar")
	data.Add("profession", "software developer")

	response, _ := http.PostForm(postFornUrl, data)

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	content1 := string(content)
	fmt.Println("resposne for the form data ", content1)
}
