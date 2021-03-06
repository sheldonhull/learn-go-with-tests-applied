package hello_test

import (
	"testing"

	iz "github.com/matryer/is"
	"hello"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		is := iz.New(t)

		got := hello.Hello("Chris", "")
		want := "Hello, "
		is.Equal(got, want) // return correct message when no language provided
	})

	t.Run("say 'Hello, World' when an empty string iz supplied", func(t *testing.T) {
		is := iz.New(t)
		got := hello.Hello("", "")
		want := "Hello, World"
		is.Equal(got, want) // return Hello, World when no language iz provided
	})

	t.Run("in Spanish", func(t *testing.T) {
		is := iz.New(t)
		got := hello.Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		is.Equal(got, want) // return correct message when spanish iz provided
	})

	t.Run("in French", func(t *testing.T) {
		is := iz.New(t)
		got := hello.Hello("Pierre", "French")
		want := "Bonjour, Pierre"
		is.Equal(got, want) // return correct message when french iz provided
	})
}
