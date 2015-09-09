package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Spawn struct {
	Url string `json:"url"`
}

// From https://www.socketloop.com/tutorials/golang-how-to-generate-random-string
func randStr() string {
	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, 12)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func spawn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s.%s", randStr(), os.Getenv("THEBE_SERVER_BASE_URL") )
		s := Spawn{ Url: url }
		dat, err := json.MarshalIndent(&s, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(dat))
	}
}
