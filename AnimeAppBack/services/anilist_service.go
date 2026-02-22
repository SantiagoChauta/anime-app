package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"AnimeApp/models"
)

const anilistURL = "https://graphql.anilist.co"

type anilistResponse struct {
	Data struct {
		Page struct {
			Media []struct {
				ID       int `json:"id"`
				Episodes int `json:"episodes"`
				Title    struct {
					Romaji string `json:"romaji"`
				} `json:"title"`
				CoverImage struct {
					Large string `json:"large"`
				} `json:"coverImage"`
			} ``
		} `json:"page"`
	} `json:"data"`
}

func SearchMedia(body map[string]any) ([]models.Anime, error) {
	raw, err := fetchFromAniList(body)
	if err != nil {
		return nil, err
	}

	decoded, err := decodeAniListResponse(raw)
	if err != nil {
		return nil, err
	}

	return mapToModels(decoded), nil
}

func fetchFromAniList(body map[string]any) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		anilistURL,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("anilist respondio con status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func decodeAniListResponse(raw []byte) (*anilistResponse, error) {
	var decoded anilistResponse

	if err := json.Unmarshal(raw, &decoded); err != nil {
		return nil, err
	}

	if decoded.Data.Page.Media == nil {
		return nil, errors.New("respuesta vacia de AniliList")
	}

	return &decoded, nil
}

func mapToModels(decoded *anilistResponse) []models.Anime {
	result := make([]models.Anime, 0, len(decoded.Data.Page.Media))

	for _, media := range decoded.Data.Page.Media {
		result = append(result, models.Anime{
			ID:       media.ID,
			Title:    media.Title.Romaji,
			Episodes: media.Episodes,
			Image:    media.CoverImage.Large,
		})
	}
	return result
}
