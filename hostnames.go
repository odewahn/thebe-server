package main

import (
  "crypto/rand"
)

// From https://www.socketloop.com/tutorials/golang-how-to-generate-random-string
func randStr() string {
	dictionary := "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, 12)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}


func generateHostName() (string) {
  return randStr()
}
