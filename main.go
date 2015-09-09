package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {

	// Load the environment variables we need
	err := godotenv.Load()
	if err != nil {
    	log.Fatal("Error loading .env file")
	}


	// Read the port
	port := os.Getenv("PORT")

	mux := mux.NewRouter()
  // mux.HandleFunc("/events", get_events(dbmap)).Methods("GET")
  // mux.HandleFunc("/events/{year}", get_events_by_year(dbmap)).Methods("GET")

  n := negroni.Classic()
  n.UseHandler(mux)
	log.Printf("Listening on port %s\n", port)
  n.Run(":" + port)

}
