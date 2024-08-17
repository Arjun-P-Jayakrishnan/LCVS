package main

import (
  "log"
  "encoding/json"
  "net/http"
)

func respondWithJSON(w http.ResponseWriter,code int,payload interface{}){
    
  dat,err := json.Marshal(payload)

  if err!=nil {

    log.Printf("Failed to marshal JSON resposnse : %v",payload)
    w.WriteHeader(500)
    
    return
  }

  w.Header().Add("Content-Type","applicatio/json")
  w.WriteHeader(code)
  w.Write(dat)

}

func respondWithError(w http.ResponseWriter,code int,msg string){

    if code >499 {

      log.Println("Responding with 5XX error: ",msg)

    }

    type errResponse struct{
        Error string `json:error`
    }
    
    respondWithJSON(w,code,errResponse{
      Error: msg,
    })
}
