package http

import (
	"net/http"

	"AnimeApp/internal/http/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/anime/latest", handlers.GetLatestAnimes)
}
