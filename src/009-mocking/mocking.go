package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}
type DefaultSleeper struct{}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, fmt.Sprintf("%d\n", i))
		sleeper.Sleep()
	}
	sleeper.Sleep()
	fmt.Fprintf(out, "Go!")
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
