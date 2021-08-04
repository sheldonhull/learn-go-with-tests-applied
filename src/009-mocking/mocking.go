package main

import (
	"bytes"
	"fmt"
)

func Countdown(out *bytes.Buffer, count int) {
	fmt.Fprintf(out, fmt.Sprintf("%d", count))
}

func main() {
}
