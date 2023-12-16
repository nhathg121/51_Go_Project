package main

import (
	"fmt"
	"log"
	"net/http"
)



func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err!= nil {
		log.Fatal(err)
	}
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {

		http.Error(w, "404 Not found", http.StatusNotFound)
		return 
	}
	if r.Method!= "GET"{
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello !")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
			fmt.Fprintf(w,"ParseForm() err: %v", err)
			return
	}
	fmt.Fprintf(w,"POST  request successful! \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name =  %s\n", name)
	fmt.Fprintf(w,"Address =  %s\n", address)
}