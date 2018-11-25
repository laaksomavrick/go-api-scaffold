package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthz", healthCheck).Methods("GET")
	router.HandleFunc("/message", handleQryMessage).Methods("GET")
	router.HandleFunc("/m/{msg}", handleURLMessage).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server is up!")
}

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func handleURLMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
