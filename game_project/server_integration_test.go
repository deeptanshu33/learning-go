package main

import (
	"fmt"
	server "game_project/http-server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrieving(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := server.PlayerServer{Store: store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}

func newPostScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", name), nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return request
}

func assertStatus(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
