package main

func findKthLargest(nums []int, k int) int {
	nums2 := make([]int, len(nums))
	for i, x := range nums {
		nums2[i] = -x
	}

	heap := buldHeap(nums2)
	for range k - 1 {
		heap.popMin()
	}
	num, _ := heap.popMin()

	return -num
}
