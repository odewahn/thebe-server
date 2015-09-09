package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
	"github.com/samalba/dockerclient"
	"crypto/tls"
)

func main() {

	// Load the environment variables we need
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the port
	port := os.Getenv("PORT")
	tls := tls.Config()


	docker, _ := dockerclient.NewDockerClient(
		os.Getenv("DOCKER_HOST"),
		*tls
	)


	mux := mux.NewRouter()
	// mux.HandleFunc("/events", get_events(dbmap)).Methods("GET")
	// mux.HandleFunc("/events/{year}", get_events_by_year(dbmap)).Methods("GET")
	mux.HandleFunc("/spawn", spawn()).Methods("GET")
	mux.HandleFunc("/list-containers", list_containers(docker)).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(mux)
	log.Printf("Listening on port %s\n", port)
	n.Run(":" + port)

}
