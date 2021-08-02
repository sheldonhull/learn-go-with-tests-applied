package di_test

import (
	"bytes"
	"di"
	"testing"

	iz "github.com/matryer/is"
)

func TestGreet(t *testing.T) {
	is := iz.New(t)
	buffer := bytes.Buffer{}
	di.Greet(&buffer, "Sheldon")
	got := buffer.String()
	want := "Hello, Sheldon"

	is.Equal(got, want) // should greet me with appropriate respect
}
