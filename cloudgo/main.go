package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // Parse (will not automatic parse unless you call it manually)
	fmt.Println("r.Form:", r.Form)  // Print out the Form to Server Side
	fmt.Println("path", r.URL.Path) // Print out the Path of the Request
	// Print out the form submitted via GET
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ", "))
	}
	// Write Response back to Client (Will be displayed on Browser)
	fmt.Fprintf(w, "Hello %+v!\n", r.Form.Get("name"))
}

func main() {
	http.HandleFunc("/", sayhelloName)       // Set Router
	err := http.ListenAndServe(":9090", nil) // Set Listening Port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
