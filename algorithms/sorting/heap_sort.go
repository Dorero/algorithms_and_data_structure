package sorting

func HeapSort(slice []int) {
	length := len(slice)

	for i := length/2 - 1; i >= 0; i-- {
		heapify(slice, length, i)
	}

	for i := length - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice, i, 0)
	}
}

func heapify(slice []int, length int, i int) {
	root := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && slice[left] > slice[root] {
		root = left
	}

	if right < length && slice[right] > slice[root] {
		root = right
	}

	if root != i {
		slice[i], slice[root] = slice[root], slice[i]

		heapify(slice, length, root)
	}
}
