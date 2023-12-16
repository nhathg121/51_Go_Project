package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"

)


type Movie struct {
	ID string `json:"id,omitempty"`  
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`

}

var movies []Movie
// func get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

// func delete movie by id
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index,item := range movies {
		if item.ID == params["id"] {
movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// func get movie by id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _,item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}


func main() {
	r := mux.NewRouter()


	movies = append(movies, Movie{ID: "1",Isbn: "111111",Title: "Movie 1",Director: &Director{FirstName: "John",LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2",Isbn: "222222",Title: "Movie 2",Director: &Director{FirstName: "John",LastName: "Doe"}})
	movies = append(movies, Movie{ID: "3",Isbn: "333333",Title: "Movie 3",Director: &Director{FirstName: "John",LastName: "Doe"}})
	movies = append(movies, Movie{ID: "4",Isbn: "444444",Title: "Movie 4",Director: &Director{FirstName: "John",LastName: "Doe"}})
	movies = append(movies, Movie{ID: "5",Isbn: "555555",Title: "Movie 5",Director: &Director{FirstName: "John",LastName: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("Get")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")



	fmt.Printf("Starting server on port 1999\n")
	log.Fatal(http.ListenAndServe(":1999", r))
}