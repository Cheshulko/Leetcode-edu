package main

func sortArray(nums []int) []int {
	heap := buldHeap(nums)

	var ans []int
	for {
		num, err := heap.popMin()
		if err != nil {
			break
		}
		ans = append(ans, num)
	}

	return ans
}
