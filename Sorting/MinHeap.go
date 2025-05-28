package main

import "errors"

type MinHeap struct {
	arr  []int
	size int
}

func buldHeap(arr []int) MinHeap {
	heap := MinHeap{
		arr:  arr,
		size: len(arr),
	}

	for i := heap.size / 2; i >= 0; i-- {
		heap.siftDown(i)
	}

	return heap
}

func (h *MinHeap) siftDown(i int) {
	for 2*i+1 < h.size {
		left := 2*i + 1
		right := 2*i + 2
		j := left

		if right < h.size && h.arr[right] < h.arr[left] {
			j = right
		}
		if h.arr[i] <= h.arr[j] {
			break
		}
		h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
		i = j
	}
}

func (h *MinHeap) siftUp(i int) {
	for h.arr[i] < h.arr[(i-1)/2] {
		h.arr[i], h.arr[(i-1)/2] = h.arr[(i-1)/2], h.arr[i]
		i = (i - 1) / 2
	}
}

func (h *MinHeap) insert(key int) {
	h.size++
	h.arr[h.size-1] = key
	h.siftUp(h.size - 1)
}

func (h *MinHeap) popMin() (int, error) {
	if h.size == 0 {
		return 0, errors.New("heap is empty")
	}

	min := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.size--
	h.siftDown(0)

	return min, nil
}
