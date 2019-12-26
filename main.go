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
	router.HandleFunc("/api/v1/users", handlers.CreateUserHandler).Methods("GET")
	//router.HandleFunc("/api/v1/auth/login", TODO).Methods("POST")
	//router.HandleFunc("/api/v1/users/{id}", TODO).Methods("PUT")
	router.HandleFunc("/api/v1/tasks", handlers.CreateTaskHandler).Methods("POST")
	//router.HandleFunc("/api/v1/tasks", TODO).Methods("GET")
	router.HandleFunc("/api/v1/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))

}
