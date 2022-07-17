package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := BinarySearch(numbers, 10)
	expected := 9

	if expected != result {
		t.Errorf("expected '%d', but got '%d'", expected, result)
	}
}

func ExampleBinarySearch() {
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2))

	// Output: 1
}