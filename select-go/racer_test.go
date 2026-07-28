package selectgo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(fastServer.URL, slowServer.URL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		slowServer.Close()
		fastServer.Close()
	})

	t.Run("return an error if a server doesn't respond withing 10 seconds", func(t *testing.T) {
		server := makeDelayedServer(11 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("Expected an error but got none")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
