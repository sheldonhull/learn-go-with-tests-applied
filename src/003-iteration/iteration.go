package iteration

import (
	"strings"
)

// Repeat will return a string repeated by the provided count.
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}
