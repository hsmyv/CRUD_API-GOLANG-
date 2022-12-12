package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github/gorilla/mux"
)

type Movie struct{
	ID 		string `json:"id"`
	Isbn 	string `json:"isbn"`
	Title 	string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

// Tenemos peliculas
var movies []Movie


//llegar a las peliculas
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

//eliminar las peliculas
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


//llegar a la pelicula
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


//crear la pelicula
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie 
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


//actualizar la pelicula
func updateMovie(w http.ResponseWriter, r *http.Request){
	//set json content type -> establecer el tipo de contenido json
	w.Header().Set("Content-Type", "application/json")
	//params -> parametros
	params := mux.Vars(r)
	
	//loop over the movies, range  -> bucle sobre las peliculas , rango
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			//delete the movie with the i.d that you've sent -> eliminar la pelicula con el i.d que has enviado
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]	
			//add a new movie- the movie that we send in the body of postman  -> anadir nueva pelicula que enviamos en postman
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}
}

//
func main(){
	//mux is the package that we use for CRUD -> mux es la paquet que usamos para CLEA
	r := mux.NewRouter()

	//We add movies to movies array  //anadimos las peliculas a los series de pelicula
	movies = append(movies, Movie{ID: "1", Isbn: "43881", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID:"2", Isbn: "45455", Title: "Movie two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	//creamos CLAE via URL -> we creat url by URL
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
