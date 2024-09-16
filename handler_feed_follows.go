package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Arjun-P-Jayakrishnan/LCVS/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedId uuid.UUID `json:"feed_idstring"`	
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
    FeedId:    params.FeedId,
	})

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't Create feed follow %v", err))
	}

  respondWithJSON(w,200,databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter,r *http.Request){

  feed,err := apiCfg.DB.GetFeeds(r.Context())
  
  if err!=nil{
    respondWithError(w,400,fmt.Sprintf("Couldn't get feeds %v",err))
    return
  }

  respondWithJSON(w,200,databaseFeedsToFeeds(feed))

 }
