package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
  "database/sql"

	"github.com/Arjun-P-Jayakrishnan/LCVS/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)




type apiConfig struct{
    DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

  dbUrl :=os.Getenv("DB_URL")
  if dbUrl == "" {
      log.Fatal("PORT is not found in the environment")
  }

  conn,err:=sql.Open("postgres",dbUrl)

  if err !=nil {
      log.Fatal("Can't connect to database:")
  }


  apiCfg :=apiConfig{
    DB:database.New(conn), 
  }

  router:=chi.NewRouter()

  router.Use(cors.Handler(cors.Options{
      AllowedOrigins: []string{"https://*","htttp://*"},
      AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
      AllowedHeaders: []string{"*"},
      ExposedHeaders: []string{"Link"},
      AllowCredentials: false,
      MaxAge:          300,

  }))

  v1Router := chi.NewRouter()
  v1Router.Get("/healthz",handlerReadiness)
  v1Router.Get("/err",handlerErr)
  v1Router.Post("/users",apiCfg.handlerCreateUser)

  router.Mount("/v1",v1Router)

  srv := &http.Server{
      Handler:router,
      Addr:":"+portString,
  }
  
  log.Printf("Server starting at Port %v",portString)
  errServer := srv.ListenAndServe()

  if errServer!=nil {
      log.Fatal(errServer)
  }

	fmt.Println("Port:", portString)
}
