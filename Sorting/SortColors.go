package main

func sortColors(nums []int) {
	var count [3]int

	for _, num := range nums {
		count[num] += 1
	}

	pos := 0
	for i := range 3 {
		for range count[i] {
			nums[pos] = i
			pos++
		}
	}
}
