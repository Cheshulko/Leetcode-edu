package main

import (
	"slices"
	"sort"
)

func heightChecker(heights []int) int {
	expected := slices.Clone(heights)
	sort.Ints(expected)

	ans := 0
	for i := range len(heights) {
		if expected[i] != heights[i] {
			ans += 1
		}
	}

	return ans
}
