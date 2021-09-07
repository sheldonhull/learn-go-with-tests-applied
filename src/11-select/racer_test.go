package racer_test

import (
	src "racer"
	"testing"

	iz "github.com/matryer/is"
)

func TestRacer(t *testing.T) {
	is := iz.New(t)

	slowURL := "https://www.facebook.com"
	fastURL := "https://www.quii.co.uk"

	want := fastURL
	got := src.Racer(slowURL, fastURL)

	is.Equal(got, want) // the fast URL should be quii obviously
}
