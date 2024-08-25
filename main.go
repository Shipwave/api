package main

import (
	"net/http"

	"github.com/campbell-frost/Shipwave/server/routes"
)

func main() {
	router := routes.SetupRoutes()
	http.ListenAndServe(":3000", router)
}
