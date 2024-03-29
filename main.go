package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "asd12312312412",
		Title: "Golang Action",
		Director: &Director{
			Firstname: "Dee",
			Lastname:  "Bee",
		},
	},
		Movie{
			ID:    "2",
			Isbn:  "ad9knasduasd",
			Title: "The Terminiator",
			Director: &Director{
				Firstname: "Jack",
				Lastname:  "Ma",
			},
		})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server listening on 8009")
	log.Fatal(http.ListenAndServe(":8009", r))

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}
