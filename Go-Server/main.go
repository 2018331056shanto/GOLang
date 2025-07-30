package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<h1>I am shanto, a Programmer</h1>")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running on : 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {

		fmt.Println("Error starting the server", err)
	}

}
