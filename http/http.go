package main

import (
	"fmt"
	"net/http"
	"strings"
)

// func ListenAndServe(addr string, handler Handler) error

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type PlayerStore interface {
	GetPlayerScore(name string) string
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(w, p.store.GetPlayerScore(player))

}

func GetPlayerScore(player string) string {
	if player == "floyd" {
		return "10"
	}

	if player == "pepper" {
		return "20"
	}

	return ""
}
