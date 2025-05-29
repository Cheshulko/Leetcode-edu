package main

import (
	"slices"
	"sort"
)

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	cnt := len(nums)

	id := make([]int, cnt)
	place := make([]int, cnt)
	for i := range cnt {
		id[i] = i
		place[i] = i
	}

	var tails [][]int

	dts := len(nums[0])
	for dl := range dts {
		sort.Slice(id, func(i, j int) bool {
			return (nums[id[i]][dts-dl-1] < nums[id[j]][dts-dl-1]) ||
				((nums[id[i]][dts-dl-1] == nums[id[j]][dts-dl-1]) && (place[id[i]] < place[id[j]]))
		})

		for i := range cnt {
			place[id[i]] = i
		}

		tails = append(tails, slices.Clone(id))
	}

	var ans []int
	for _, query := range queries {
		ans = append(ans, tails[query[1]-1][query[0]-1])
	}

	return ans
}
