package racer

import (
	"fmt"
	http "net/http"
	"time"
)

// timeoutSec
var timeoutSec = 10 * time.Second

// Racer returns the fastest responding website.
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, timeoutSec)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout waiting for %s and %s", a, b)
	}
}

// ping returns a channel that will receive a value once the url is available.
// From the Learn Tests With Go - Select Section:
// > Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
