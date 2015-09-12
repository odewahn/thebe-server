package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/samalba/dockerclient"
	"log"
	"os"
)

// Need a function to define the tls Config
//   https://github.com/ehazlett/interlock/blob/master/interlock/main.go#L14-L32

func main() {

	// Load the environment variables we need
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the port
	port := os.Getenv("PORT")

	tlsConfig, err := getTLSConfig(os.Getenv("SWARM_CREDS_DIR"))
	if err != nil {
		log.Fatal("Could not create TLS certificate.")
	}

	docker, _ := dockerclient.NewDockerClient(os.Getenv("DOCKER_HOST"), tlsConfig)

	mux := mux.NewRouter()
	// mux.HandleFunc("/events", get_events(dbmap)).Methods("GET")
	// mux.HandleFunc("/events/{year}", get_events_by_year(dbmap)).Methods("GET")
	mux.HandleFunc("/spawn", spawn(docker)).Methods("GET")
	mux.HandleFunc("/list-containers", list_containers(docker)).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(mux)
	log.Printf("Listening on port %s\n", port)
	n.Run(":" + port)

}
