package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server on port 10000\n")
	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal(err)
	}
}

// helloHandler is a handler function for the /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello !")
}

// formHandler is a handler function for the /form route
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST  request successful! \n")
	fmt.Fprintf(w, "Name: %s \n", r.FormValue("name"))
	fmt.Fprintf(w, "address: %s \n", r.FormValue("address"))

}
