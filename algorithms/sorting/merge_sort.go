package sorting

func MergeSort(slice []int) []int {
	length := len(slice)

	if length == 1 {
		return slice
	}

	mid := length / 2

	left := MergeSort(slice[:mid])
	right := MergeSort(slice[mid:])

	return mergeTwoList(left, right)
}

func mergeTwoList(first []int, second []int) []int {
	left := len(first)
	right := len(second)
	i, j := 0, 0
	result := []int{}

	for i < left && j < right {
		if first[i] < second[j] {
			result = append(result, first[i])
			i++
		} else {
			result = append(result, second[j])
			j++
		}
	}

	for i < left {
		result = append(result, first[i])
		i++
	}

	for j < right {
		result = append(result, second[j])
		j++
	}

	return result
}
