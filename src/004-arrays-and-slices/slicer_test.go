package slicer_test

import (
	"slicer"
	"testing"

	"github.com/matryer/is"
)

func TestSum(t *testing.T) {
	is := is.New(t)
	got := slicer.Sum([]int{1, 1, 1, 1, 1})
	want := 5

	is.Equal(got, want) // sum should return 3
}
