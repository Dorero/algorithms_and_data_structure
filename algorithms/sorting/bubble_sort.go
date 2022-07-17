package sorting

func BubbleSort(numbers []int) []int {
	length := len(numbers)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if numbers[j+1] < numbers[j] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}
