package router

import (
	"Amazon/user/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", handler.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handler.Getuser).Methods("GET")
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users{id}", handler.UpadateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")
	return r

}
