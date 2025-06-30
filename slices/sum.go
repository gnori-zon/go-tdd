package slices

func Sum(elements []int) int {
	sum := func(result, item int) int { return result + item }
	return Reduce(elements, sum, 0)
}

func SumAll(slices [][]int) []int {
	sumAll := func(result, item []int) []int {
		return append(result, Sum(item))
	}
	return Reduce(slices, sumAll, []int{})
}

func SumAllTails(slices [][]int) []int {
	sumTail := func(result, item []int) []int {
		if len(item) < 1 {
			return append(result, 0)
		} else {
			tail := item[1:]
			return append(result, Sum(tail))
		}
	}
	return Reduce(slices, sumTail, []int{})
}
