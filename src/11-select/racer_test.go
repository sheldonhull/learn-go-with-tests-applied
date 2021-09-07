package racer_test

import (
	http "net/http"
	"net/http/httptest"
	src "racer"
	"testing"
	"time"

	iz "github.com/matryer/is"
)

func TestRacer(t *testing.T) {
	is := iz.New(t)

	slowServer := makeDelayedServer(20 * time.Millisecond)
	defer slowServer.Close()

	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := src.Racer(slowURL, fastURL)

	is.Equal(got, want) // the fast URL should be quii obviously
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
