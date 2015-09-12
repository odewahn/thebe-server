package main

import (
	"encoding/json"
	"fmt"
	"github.com/samalba/dockerclient"
	"log"
	"net/http"
)

func list_containers(docker *dockerclient.DockerClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get only running containers
		containers, err := docker.ListContainers(false, false, "")
		if err != nil {
			log.Fatal(err)
		}

		dat, err := json.MarshalIndent(&containers, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(dat))
	}
}
