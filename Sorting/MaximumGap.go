package main

import "slices"

func maximumGap(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	ma, mi := slices.Max(nums), slices.Min(nums)

	// mi ... x1 . x2 .... x3 ... ma
	max_possible := max(1, (ma-mi)/(n-1))

	// [[[[mi .... 0] .... 1] .... 2] ..... ma 3]
	// Ex.
	// n=5
	// mi=0 ma=100
	// max_possible=(100 - 0) / (5 - 1) = 25
	// 0........1........2........3.......4

	// 0 25 50 75 100
	groups := make([][]int, 1+(ma-mi)/max_possible)
	for _, num := range nums {
		group := (num - mi) / max_possible

		if len(groups[group]) == 0 {
			groups[group] = []int{num, num}
		} else {
			groups[group][0] = min(groups[group][0], num)
			groups[group][1] = max(groups[group][1], num)
		}
	}

	ans := 0
	prev_ma := mi
	for i := range len(groups) {
		if len(groups[i]) > 0 {
			ans = max(ans, groups[i][0]-prev_ma)
			prev_ma = groups[i][1]
		}
	}

	return ans
}
