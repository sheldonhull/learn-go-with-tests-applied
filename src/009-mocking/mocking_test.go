package main //nolint: testpackage

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestCountdown(t *testing.T) {
	is := is.New(t)

	buffer := &bytes.Buffer{}
	Countdown(buffer)

	is.Equal(buffer.String(), "3\n"+"2\n"+"1\n"+"Go!") // countdown returns expected value
}
