package handler

import (
	"Amazon/user/internal/controller"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	controller.ListUsers(w, r)
}

func Getuser(w http.ResponseWriter, r *http.Request) {
	controller.Getuser(w, r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	controller.CreateUser(w, r)
}

func UpadateUser(w http.ResponseWriter, r *http.Request) {
	controller.UpadateUser(w, r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	controller.DeleteUser(w, r)
}
