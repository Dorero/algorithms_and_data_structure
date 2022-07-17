package sorting

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	slice := []int{22, 33, 1, 5, 2, 7, 9, 3}
	result := BubbleSort(slice)
	expected := []int{1, 2, 3, 5, 7, 9, 22, 33}

	if !comparableSlices(expected, result) {
		t.Errorf("expected '%v, but got '%v'", expected, result)
	}
}

func ExampleBubbleSort() {
	slice := []int{6, 3, 8, 2, 1, 5}
	result := SelectionSort(slice)

	fmt.Println(result)

	// Output: [1 2 3 5 6 8]
}
