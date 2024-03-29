package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Consumer struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

var consumers []Consumer

func getConsumer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(consumers)
}

func getConsumers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range consumers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Consumer{})
}

func createConsumer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var consumer Consumer
	_ = json.NewDecoder(r.Body).Decode(&consumer)
	consumer.ID = params["id"]
	consumers = append(consumers, consumer)
	json.NewEncoder(w).Encode(consumers)
}

func deleteConsumer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range consumers {
		if item.ID == params["id"] {
			consumers = append(consumers[:index], consumers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(consumers)
}

func main() {
	router := mux.NewRouter()

	consumers = append(consumers, Consumer{ID: "1", Username: "user1", Email: "example@example.com"})
	consumers = append(consumers, Consumer{ID: "2", Username: "user2", Email: "example_two@example.com"})

	router.HandleFunc("/consumers", getConsumer).Methods("GET")
	router.HandleFunc("/consumers/{id}", getConsumers).Methods("GET")
	router.HandleFunc("/consumers/{id}", createConsumer).Methods("POST")
	router.HandleFunc("/consumers/{id}", deleteConsumer).Methods("DELETE")

	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
