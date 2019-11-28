package main

import (
	"github.com/gorilla/mux"
	"github.com/project-management-api/handlers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/api/v1/users", handlers.CreateUserEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
