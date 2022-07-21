package main 

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"ID"`
	Isbn string `json:"Isbn"`
	Title string `json:"Title"`
	Director *Director `json:"Director"`
}

type Director struct {
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["ID"]{
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params ["ID"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["ID"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn:"438227", Title:"Movie One", Director : &Director{FirstName:"Rishi", LastName:"Yadav"}})
	movies = append(movies, Movie{ID: "2", Isbn:"45455", Title:"Movie Two", Director : &Director{FirstName:"Siddharth", LastName:"Yadav"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}