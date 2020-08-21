package main

import (
	"net/http"

	"github.com/scys12/simple-api-go/routes"
)

func main() {
	r := routes.NewRouter()

	http.ListenAndServe(":8080", r)
}
