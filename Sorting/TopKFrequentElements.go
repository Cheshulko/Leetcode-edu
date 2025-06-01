package main

import (
	"container/heap"
)

type Item struct {
	Count int
	Num   int
}

type ItemHeap []*Item

func (h *ItemHeap) Len() int {
	return len(*h)
}

func (h *ItemHeap) Less(i, j int) bool {
	return (*h)[i].Count > (*h)[j].Count
}

func (h *ItemHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]

	return item
}

func (h ItemHeap) Peek() *Item {
	if len(h) == 0 {
		return nil
	}

	return h[0]
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	items := &ItemHeap{}
	for num, freq := range count {
		*items = append(*items, &Item{Num: num, Count: freq})

	}

	heap.Init(items)

	result := make([]int, 0, k)
	for i := 0; i < k && items.Len() > 0; i++ {
		item := heap.Pop(items).(*Item)
		result = append(result, item.Num)
	}

	return result
}
