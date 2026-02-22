package handlers

import (
	"encoding/json"
	"net/http"

	"AnimeApp/services"
)

func GetLatestAnimes(w http.ResponseWriter, r *http.Request) {
	body := map[string]any{
		"query": `
			query {
  			Page(page: 1, perPage: 20) {
    			media(
      			type: ANIME
      			sort: START_DATE_DESC
		      ){
      			id
      			title {
        			romaji
      			}
      			episodes
      			coverImage {
        			large
      			}
      			startDate {
        			year
        			month
        			day
      			}
    			}
  			}
			}
	`,
	}
	animes, err := services.SearchMedia(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(animes); err != nil {
		http.Error(w, "error al escribir la respuesta", http.StatusInternalServerError)
		return
	}
}
