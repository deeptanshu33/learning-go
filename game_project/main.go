package main

import (
	server "game_project/http-server"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	server := &server.PlayerServer{Store: NewInMemoryPlayerStore()}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
