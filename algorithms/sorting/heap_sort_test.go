package sorting

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	slice := []int{22, 33, 1, 5, 2, 7, 9, 3}
	HeapSort(slice)
	expected := []int{1, 2, 3, 5, 7, 9, 22, 33}

	if !comparableSlices(expected, slice) {
		t.Errorf("expected '%v, but got '%v'", expected, slice)
	}
}

func ExampleHeapSort() {
	slice := []int{6, 3, 8, 2, 1, 5}
	HeapSort(slice)

	fmt.Println(slice)

	// Output: [1 2 3 5 6 8]
}
