package main

func mergeSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}

	left := mergeSort(nums[:l/2])
	right := mergeSort(nums[l/2:])

	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func sortArray(nums []int) []int {
	return mergeSort(nums)
}
