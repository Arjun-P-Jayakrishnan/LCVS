package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Arjun-P-Jayakrishnan/LCVS/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter,r *http.Request){

    //the json format we want the data to be in.
    type parameters struct{
      Name string `json:"name"`
    }

    //parse the data in request to json.
    decoder := json.NewDecoder(r.Body)

    //create a instance of the struct to have a storage space.
    params  := parameters{}

    //Marshall the data in json to the required format and store in the location given as pointer.
    err := decoder.Decode(&params)

    //user facing issue
    if err!=nil {
          respondWithError(w,400,fmt.Sprintf("Error parsing JSON: %v",err))
          return
    }

    // create user using http Context, 
    user,err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
        ID:uuid.New(),
        CreatedAt:time.Now().UTC(),
        UpdatedAt:time.Now().UTC(),
        Name: params.Name,
    })

    if err!=nil {
      respondWithError(w,400,fmt.Sprintf("Couldnt create user: %v",err))
      return
    }

    respondWithJSON(w,201,databaseUserToUser(user))
}

func (apiCfg *apiConfig)handlerGetUser(w http.ResponseWriter,r *http.Request, user database.User){
  
  
  respondWithJSON(w,200,databaseUserToUser(user))
}


