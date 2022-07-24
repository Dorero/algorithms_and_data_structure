package sorting

func QuickSort(slice []int) []int {
	length := len(slice)

	if length <= 1 {
		return slice
	}

	elem := slice[length/2]

	less := filter(func(e int) bool { return e < elem }, slice)
	mid := filter(func(e int) bool { return e == elem }, slice)
	more := filter(func(e int) bool { return e > elem }, slice)

	return append(append(QuickSort(less), mid...), QuickSort(more)...)
}

func filter(fn func(first int) bool, slice []int) []int {
	res := []int{}

	for _, value := range slice {
		if fn(value) {
			res = append(res, value)
		}
	}

	return res
}
