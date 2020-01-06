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
	router.HandleFunc("/api/v1/users", handlers.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/api/v1/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/api/v1/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	//TODO: router.HandleFunc("/api/v1/auth/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/tasks", handlers.CreateTaskHandler).Methods("POST")
	//TODO: router.HandleFunc("/api/v1/tasks/{project-id}", handlers.GetAllTasksHandler).Methods("GET")
	//TODO: router.HandleFunc("/api/v1/tasks/{id}", handlers.GetTaskHandler).Methods("GET")
	//TODO: router.HandleFunc("/api/v1/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT")
	//TODO: router.HandleFunc("/api/v1/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
