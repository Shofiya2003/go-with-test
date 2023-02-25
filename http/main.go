package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}
	//the server might return an error for that reason we wrap it into a log.Fata
	log.Fatal(http.ListenAndServe(":5000", server))
}
