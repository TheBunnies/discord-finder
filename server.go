package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TheBunnies/discord_finder/discord"
	"github.com/gorilla/mux"
)

var (
	token, _ = loadConfig()
)

type Token struct {
	Body string `json:"token"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	user, err := discord.GetUser(id, token)
	if err != nil {
		message := ErrorMessage{Message: "An error occured while communicating to discord API."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func setupRoutes() {
	r := mux.NewRouter()
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/api/{id}", GetUser).Methods("GET")
	r.Use(mux.CORSMethodMiddleware(r))
	log.Fatal(srv.ListenAndServe())
}

func loadConfig() (string, error) {
	if _, err := os.Stat("config.json"); err != nil {
		os.Create("config.json")
		file, err := os.OpenFile("config.json", os.O_APPEND, os.ModeAppend)
		if err != nil {
			return "", err
		}
		defer file.Close()
		token := Token{
			Body: "your token goes here",
		}
		err = json.NewEncoder(file).Encode(token)
		if err != nil {
			return "", err
		}
		return "", err
	}
	file, err := os.Open("config.json")
	if err != nil {
		return "", err
	}
	defer file.Close()
	token := Token{}
	json.NewDecoder(file).Decode(&token)
	return token.Body, nil
}

func main() {
	log.Print("Everything worked just fine!")
	setupRoutes()
}
