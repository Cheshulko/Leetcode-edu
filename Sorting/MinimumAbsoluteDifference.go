package main

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	diff := 2_000_000 + 1
	for i := range len(arr) - 1 {
		diff = min(diff, arr[i+1]-arr[i])

	}

	var ans [][]int

	nums := make(map[int][]int)
	for i, x := range arr {
		need := x - diff

		for _, j := range nums[need] {
			ans = append(ans, []int{arr[j], x})
		}

		nums[x] = append(nums[x], i)
	}

	return ans
}
