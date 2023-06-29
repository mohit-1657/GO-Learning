package main

import (
	"encoding/json"
	"fmt"
)

// type is the base interface for all the data type in go
type course struct {
	//Name     string // normal way which will gives you the same name

	// if want to change the varible name into the json response according to our need we can do that
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`              // (-) this is used to hide the varibles from the api consumption // sp paasword will not be visible  into api in this
	Tag      []string `json:"tags,omitempty"` // this will not show if the value is getting null  // but one more thing it keep in mind with space it can give some error diue to that
}

func main() {

	fmt.Println("Lets start with the json manipulation")

	// fmt.Println("Encoding part of json start from here")
	// EncodeJson()

	fmt.Println("Decoding part of json start from here")
	Decodejson()
}

func EncodeJson() {

	listOfCoursec := []course{
		{"React-js", 2000, "Online.go.react", "etggeeg", []string{"start react", "build website"}},
		{"NODE-js", 3000, "Online.go,node", "etggeecbg", []string{"start node", "build website"}},
		{"Mern", 2000, "Online.go.mern", "edbvbzdsybdy", nil},
	}

	// now packaging this data into json data

	// json.marhsel is used to do encoding and the decoding of the json data

	//finaljson, err := json.Marshal(listOfCoursec)

	//if we want the result in indentent format then
	finaljson, err := json.MarshalIndent(listOfCoursec, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finaljson)

}

func Decodejson() {

	jsonFromTheWeb := []byte(`

			{
                "coursename": "Reactjs",
                "Price": 2000,
                "Platform": "Online.go.react",
                "Password": "etggeeg",
                "tags": ["start react","build website"]
        	}
		`)

	// verifying the data is correct json or not

	var coursesList course

	checkValid := json.Valid(jsonFromTheWeb)

	if checkValid {

		fmt.Println("this is a valid json format")
		json.Unmarshal(jsonFromTheWeb, &coursesList)

		//	fmt.Println("", coursesList)
		fmt.Printf("%#v\n", coursesList) // %#v is used
	} else {
		fmt.Println("json was not valid")

	}

	// in some cases you want add the data

	var myOnlineData map[string]interface{}
	// so here we know in json that first key always will be a string and values could be anything int string array
	// map deca.artion is key is string and for values it is interface
	fmt.Println("onlie json changed")
	json.Unmarshal(jsonFromTheWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	// now iterartring through the loop

	for k, v := range myOnlineData {
		fmt.Printf("key :%v value %v \n", k, v)

	}

}
