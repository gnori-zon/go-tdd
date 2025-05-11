package slices

func Sum(elements []int) int {
	sum := 0
	for _, element := range elements {
		sum += element
	}
	return sum
}

func SumAll(slices [][]int) []int {
	sums := make([]int, len(slices))
	for idx, slice := range slices {
		sums[idx] = Sum(slice)
	}
	return sums
}

func SumAllTails(slices [][]int) []int {
	sums := make([]int, len(slices))
	for idx, slice := range slices {
		if len(slice) < 1 {
			sums[idx] = 0
		} else {
			tail := slice[1:]
			sums[idx] = Sum(tail)
		}
	}
	return sums
}
