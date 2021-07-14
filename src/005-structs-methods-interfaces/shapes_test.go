package shapes_test

import (
	"testing"

	"shapes"
)

func TestPerimeter(t *testing.T) {
	t.Run("test a square", func(t *testing.T) {
		is := is.New(t)
		got := shapes.Perimeter(10.0, 10.0)
		want := 40.0
		is.Equal(got, want) // calculate perimeter
	})
	t.Run("test a rectangle", func(t *testing.T) {
		is := is.New(t)
		got := shapes.Perimeter(4.0, 8.0)
		want := 24.0
		is.Equal(got, want) // calculate perimeter
	})
}

func TestArea(t *testing.T) {
	t.Run("calculate area of a square", func(t *testing.T) {
		is := is.New(t)
		got := shapes.Area(10.0, 10.0)
		want := 100.0
		is.Equal(got, want)
	})
}
