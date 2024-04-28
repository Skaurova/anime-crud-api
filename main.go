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

type Anime struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Episodes string    `json:"episodes"`
	Studio   string    `json:"studio"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var animes []Anime

func getAnimes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animes)
}

func deleteAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range animes {
		if item.ID == params["id"] {
			animes = append(animes[:index], animes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(animes)

}

func getAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range animes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var anime Anime
	_ = json.NewDecoder(r.Body).Decode(&anime)
	anime.ID = strconv.Itoa(rand.Intn(100000000))
	animes = append(animes, anime)
	json.NewEncoder(w).Encode(anime)

}

func updateAnime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range animes {
		if item.ID == params["id"] {
			animes = append(animes[:index], animes[index+1:]...)
			var anime Anime
			_ = json.NewDecoder(r.Body).Decode(&anime)
			anime.ID = params["id"]
			animes = append(animes, anime)
			json.NewEncoder(w).Encode(anime)
		}
	}
}

func main() {
	r := mux.NewRouter()

	animes = append(animes, Anime{ID: "1", Title: "Fullmetal Alchemist 2003", Episodes: "51", Studio: "Bones", Director: &Director{Firstname: "Seiji", Lastname: "Mizushima"}})
	animes = append(animes, Anime{ID: "2", Title: "Mawaru Penguindrum", Episodes: "24", Studio: "Brain's Base", Director: &Director{Firstname: "Kunihiko", Lastname: "Ikuhara"}})

	r.HandleFunc("/animes", getAnimes).Methods("GET")
	r.HandleFunc("/animes/{id}", getAnime).Methods("GET")
	r.HandleFunc("/animes", createAnime).Methods("POST")
	r.HandleFunc("/animes/{id}", updateAnime).Methods("PUT")
	r.HandleFunc("/animes/{id}", deleteAnime).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
