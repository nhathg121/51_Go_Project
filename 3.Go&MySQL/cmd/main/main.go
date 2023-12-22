package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	fmt.Printf("%s", r)
	fmt.Println(" ")
	routes.RegisterBookStoreRoutes(r)
	fmt.Println(123)

	//http.Handle("/", r)
	fmt.Printf("Starting server on port 9010\n")
	log.Fatal(http.ListenAndServe(":9010", r))
}
