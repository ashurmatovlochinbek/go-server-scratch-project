package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rssagg/internal/database"
	"time"
	"github.com/google/uuid"
)


func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}


	feed, error := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})

	if error != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %s", error))
		return
	}

	responseWithJSON(w, 201, databaseFeedtoFeed(feed))
}	


func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %s", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedstoFeeds(feeds))
}

