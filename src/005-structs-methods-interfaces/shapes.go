package shapes

// Perimeter returns an int by calculating a given height and width as inputs.
//nolint: gomnd
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func Area(width float64, height float64) float64 {
	return width * height
}
