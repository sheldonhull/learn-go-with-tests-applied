package main

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, fmt.Sprintf("%d\n", i))
	}
	fmt.Fprintf(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
