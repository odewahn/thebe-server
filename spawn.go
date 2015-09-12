package main

import (
	"encoding/json"
	"fmt"
	"github.com/samalba/dockerclient"
	"log"
	"net/http"
	"os"
)

type Spawn struct {
	Url string `json:"url"`
}

func spawn(docker *dockerclient.DockerClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		hostname := generateHostName()
		err := launchNotebook(docker, hostname, "ipython/scipystack")

		url := fmt.Sprintf("%s.%s", hostname, os.Getenv("THEBE_SERVER_BASE_URL"))
		s := Spawn{Url: url}
		dat, err := json.MarshalIndent(&s, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(dat))
	}
}
