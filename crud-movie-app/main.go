package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Idbn string `json:"idbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}


var movies[] Movie

func getMovies(w http.ResponseWriter,r *http.Request){
      w.Header().Set("content-type","application/json")
	  json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","application/json")
	//wants id to delete
	params := mux.Vars(r)
	for index,item := range movies {
        if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,  r *http.Request){
	w.Header().Set("content-type","application/json")
	// same as del wants req id to be deleted
	params := mux.Vars(r)
	for _,item := range movies {
		if item.ID == params["id"]{
            json.NewEncoder(w).Encode(item)
			return
		}
	}

}


func createMovie(w http.ResponseWriter, r * http.Request){
	w.Header().Set("content/type","application/json")
	var movie Movie   // variable to create new movie 
	_=json.NewDecoder(r.Body).Decode(&movie)   // decode the movie 
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie) // insert that movie to our slice --> movies 
	json.NewEncoder(w).Encode(movie)

}


func updateMovie(w http.ResponseWriter, r * http.Request){
	// delete + create 
    w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	for index,item := range movies {
        if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
            var movie Movie   // variable to create new movie 
	        _=json.NewDecoder(r.Body).Decode(&movie)   // decode the movie 
	        movie.ID = params["id"]
	        movies = append(movies, movie) 
            json.NewEncoder(w).Encode(movie)
			return 
		}
}
}


func main(){

	r := mux.NewRouter()
	
	movies = append(movies,Movie{ID:"1",Idbn:"23456",Title:"samshera",Director: &Director{Firstname: "Ayan",Lastname: "Mukherjee"}})
	movies = append(movies,Movie{ID:"2",Idbn:"23457",Title:"shershah",Director: &Director{Firstname: "Sid",Lastname: "Malhotra"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")


	fmt.Printf("Server is running at port 5000\n")
	log.Fatal(http.ListenAndServe(":5000",r))
}