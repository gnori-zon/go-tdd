package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// output: 3
}

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expectedSum := 4

	if sum != expectedSum {
		t.Errorf("want: '%d' but got: '%d'", expectedSum, sum)
	}
}
