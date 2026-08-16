package main

import (
	"encoding/json"
	"fmt"
	server "game_project/http-server"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRecordingWinsAndRetrieving(t *testing.T) {
	store := NewInMemoryPlayerStore()
	httpserver := server.NewPlayerServer(store)
	player := "Pepper"

	httpserver.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	httpserver.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	httpserver.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	t.Run("get player score", func(t *testing.T) {
		response := httptest.NewRecorder()
		httpserver.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("returns league table", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)

		httpserver.ServeHTTP(response, request)
		want := []server.Player{
			{Name: "Pepper", Wins: 3},
		}

		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, want)
	})
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []server.Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t testing.TB, got, leagueTable []server.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, leagueTable) {
		t.Errorf("got %v, want %v", got, leagueTable)
	}
}

func newPostScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
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
