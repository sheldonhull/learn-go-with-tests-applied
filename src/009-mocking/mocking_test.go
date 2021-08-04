package main //nolint: testpackage

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestCountdown(t *testing.T) {
	is := is.New(t)

	buffer := &bytes.Buffer{}
	Countdown(buffer, 3)

	is.Equal(buffer.String(), "3") // countdown returns expected value
}
