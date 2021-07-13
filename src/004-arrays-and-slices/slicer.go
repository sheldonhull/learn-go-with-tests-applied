package slicer

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(intSlices ...[]int) (sums []int) {
	for _, s := range intSlices {
		sums = append(sums, Sum(s))
	}

	return sums
}
