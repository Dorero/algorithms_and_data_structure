package sorting

func SelectionSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		min := i

		for j := i + 1; j < length; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}

	return nums
}
