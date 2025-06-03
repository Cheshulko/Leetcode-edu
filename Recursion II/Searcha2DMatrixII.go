package main

func searchMatrix(matrix [][]int, target int) bool {
	var search func(matrix [][]int, target int, is int, ie int, js int, je int) bool
	search = func(matrix [][]int, target int, is int, ie int, js int, je int) bool {
		li, lj := ie-is, je-js
		if li < 0 || lj < 0 {
			return false
		}
		if li <= 1 && lj <= 1 {
			return matrix[is][js] == target ||
				matrix[is][je] == target ||
				matrix[ie][js] == target ||
				matrix[ie][je] == target
		}

		im, jm := is+(li/2), js+(lj/2)

		if target <= matrix[im][jm] {
			return search(matrix, target, is, im, js, jm) ||
				search(matrix, target, im, ie, js, jm) ||
				search(matrix, target, is, im, jm, je)
		} else {
			return search(matrix, target, im, ie, jm, je) ||
				search(matrix, target, im, ie, js, jm) ||
				search(matrix, target, is, im, jm, je)
		}
	}

	n, m := len(matrix), len(matrix[0])

	return search(matrix, target, 0, n-1, 0, m-1)
}
