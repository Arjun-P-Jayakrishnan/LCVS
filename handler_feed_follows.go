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

	feedFollows, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
    FeedID:    params.FeedId,
	})

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't Create feed follow %v", err))
	}

  respondWithJSON(w,200,databaseFeedFollowToFeedFollow(feedFollows))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter,r *http.Request,user database.User){ 

  feed,err := apiCfg.DB.GetFeedFollows(r.Context(),user.ID) 

if err!=nil{ 
    respondWithError(w,400,fmt.Sprintf("Couldn't get feed follows %v",err))
    return 
  } 
  respondWithJSON(w,201,databaseFeedFollowsToFeedFollows(feed))
}
