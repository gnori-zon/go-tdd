package slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum should return sum for all elements from slice", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5}
		sum := Sum(elements)
		expected := 15
		assertSumForElements(t, expected, sum, elements)
	})
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func TestSumAll(t *testing.T) {
	t.Run("sum should return sums for each slice from slice", func(t *testing.T) {
		elementsSlices := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
		expected := []int{15, 40, 65}
		sum := SumAll(elementsSlices)
		assertAllSumsForElements(t, expected, sum, elementsSlices)
	})
}
func BenchmarkSumAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumAll([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}})
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("sumAllTails should return sum for each slices(without 1 element) from slice", func(t *testing.T) {
		elementsSlices := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
		expected := []int{14, 34, 54}
		sum := SumAllTails(elementsSlices)
		assertAllSumsForElements(t, expected, sum, elementsSlices)
	})

	t.Run("sumAllTest should safety for empty slice", func(t *testing.T) {
		elementsSlices := [][]int{{}}
		expected := []int{0}
		sum := SumAllTails(elementsSlices)
		assertAllSumsForElements(t, expected, sum, elementsSlices)
	})
}

func assertSumForElements(t testing.TB, expectedSum int, actualSum int, elements []int) {
	t.Helper()
	if actualSum != expectedSum {
		t.Errorf("Expected %v but got %v for slices: %v", expectedSum, actualSum, elements)
	}
}

func assertAllSumsForElements(t testing.TB, expectedSum []int, actualSum []int, elements [][]int) {
	t.Helper()
	if !slices.Equal(actualSum, expectedSum) {
		t.Errorf("Expected %v but got %v for slices: %v", expectedSum, actualSum, elements)
	}
}
