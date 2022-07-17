package search

func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		var mid = (low + high) / 2

		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
