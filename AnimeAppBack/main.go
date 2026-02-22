package main

import (
	"log"
	"net/http"

	mihttp "AnimeApp/internal/http"
	"AnimeApp/internal/http/middleware"
)

func main() {
	mux := http.NewServeMux()
	mihttp.RegisterRoutes(mux)
	handler := middleware.WithCors(mux)

	log.Println("Servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
