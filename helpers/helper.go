package helpers

func comparableSlices(numbers []int, elements []int) bool {
	if len(numbers) != len(elements) {
		return false
	}

	for index, value := range numbers {
		if value != elements[index] {
			return false
		}
	}

	return true
}
