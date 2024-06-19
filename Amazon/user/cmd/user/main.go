package main

import (
	"Amazon/user/internal/router"
	"net/http"
)

func main() {
	r := router.NewRouter()

	http.ListenAndServe(":8080", r)
}
