package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	server := &PlayerServer{Store: &store}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("get floyd's score", func(t *testing.T) {
		request := NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("missing player returns 404", func(t *testing.T) {
		request := NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{}, winCalls: make([]string, 0)}
	server := &PlayerServer{&store}

	t.Run("records win on post", func(t *testing.T) {
		request := NewPostScoreRequest("Pepper")
		response := httptest.NewRecorder()
		player := "Pepper"

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls on RecordWin, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("Expected player %q, got %q", player, store.winCalls[0])
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := &PlayerServer{&store}

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
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

func NewGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func NewPostScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}
