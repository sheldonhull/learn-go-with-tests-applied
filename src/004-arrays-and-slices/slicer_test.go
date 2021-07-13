package slicer_test

import (
	"testing"

	"slicer"

	"github.com/matryer/is"
)

func TestSum(t *testing.T) {
	is := is.New(t)
	got := slicer.Sum([5]int{1,1,1,1,1})
	want := 5

	is.Equal(got,want) // sum should return 3
}
