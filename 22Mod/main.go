package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	greeter()
	r := mux.NewRouter()

	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {

	fmt.Println("Hey are you coming to use mod ")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> this is a you tube channel to learn for frere </h1>"))
}
