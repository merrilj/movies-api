package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"
	
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	
)

func AllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not yet implemented")
}

func FindMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not yet implemented")
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not yet implemented")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not yet implemented")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMovies).Methods("GET")
	r.HandleFunc("/movies", FindMovie).Methods("POST")
	r.HandleFunc("/movies", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovie).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}