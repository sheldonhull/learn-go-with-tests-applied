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
	t.Run("when one server is faster", func(t *testing.T) {
		is := iz.New(t)
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := src.Racer(slowURL, fastURL)

		is.NoErr(err)       // no timeout error should occur
		is.Equal(got, want) // should match fastURL
	})

	t.Run("with both servers timeout an error is returned", func(t *testing.T) {
		is := iz.New(t)
		serverA := makeDelayedServer(11 * time.Second)
		defer serverA.Close()

		serverB := makeDelayedServer(12 * time.Second)
		defer serverB.Close()

		serverAURL := serverA.URL
		serverBURL := serverB.URL

		want := ""
		got, err := src.Racer(serverAURL, serverBURL)
		is.True(err != nil) // an error should be returned
		is.Equal(got, want) // no url should be returned due to timeout
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
