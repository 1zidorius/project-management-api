package main

import (
	"github.com/1zidorius/project-management-api/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Index).Methods("GET")
	router.HandleFunc("/api/v1/users", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/api/v1/task", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/v1/task", handlers.UpdateTaskHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))

}
