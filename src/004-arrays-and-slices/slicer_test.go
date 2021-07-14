package slicer_test

import (
	"testing"

	"slicer"
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

func TestSumAll(t *testing.T) {
	t.Run("passing 1 slice", func(t *testing.T) {
		is := is.New(t)
		got := slicer.SumAll([]int{1, 3, 1, 1, 1})
		want := []int{5}
		is.Equal(got, want) // 1 slice should be returned with sum of 5
	})
	t.Run("passing 2 slices", func(t *testing.T) {
		is := is.New(t)
		got := slicer.SumAll([]int{5, 5, 5}, []int{10, 10, 10})
		want := []int{15, 30}
		is.Equal(got, want) // 2 slices should be returned containing the sum of the 2 slices
	})
}
