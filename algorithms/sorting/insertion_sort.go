package sorting

func InsertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		j := i - 1
		key := numbers[i]

		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j--
		}

		numbers[j+1] = key
	}

	return numbers
}
