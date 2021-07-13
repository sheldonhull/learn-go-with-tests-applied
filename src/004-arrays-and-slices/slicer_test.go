package slicer_test

import (
	"slicer"
	"testing"

	"github.com/matryer/is"
)

func TestSum(t *testing.T) {
	is := is.New(t)

	t.Run("input size of 5", func(t *testing.T) {
		got := slicer.Sum([]int{1, 1, 1, 1, 1})
		want := 5
		is.Equal(got, want) // input should be summed
	})
	t.Run("input size of 3", func(t *testing.T) {
		got := slicer.Sum([]int{5, 5, 5})
		want := 15
		is.Equal(got, want) // input should be summed
	})
}
