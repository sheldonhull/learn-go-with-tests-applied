package main //nolint: testpackage

import (
	"bytes"
	"testing"

	iz "github.com/matryer/is"
)

func TestCountdown(t *testing.T) {
	is := iz.New(t)

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)

	is.Equal(buffer.String(), "3\n"+"2\n"+"1\n"+"Go!") // countdown returns expected value
	is.Equal(4, spySleeper.Calls)                      // spySleeper equals

	t.Run("sleep before every print", func(t *testing.T) {
		is := iz.New(t)
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		is.Equal(spySleepPrinter.Calls, want) // spySleepPrinter shows correct order of calls
	})
}
