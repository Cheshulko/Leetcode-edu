package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var check func(*TreeNode) (int, int, bool)
	check = func(node *TreeNode) (int, int, bool) {
		result := true
		mi, ma := node.Val, node.Val
		if node.Left != nil {
			leftMi, leftMa, leftResult := check(node.Left)
			result = result && leftResult
			result = result && (leftMa < node.Val)
			mi = min(mi, leftMi)
		}
		if node.Right != nil {
			rightMi, rightMa, rightResult := check(node.Right)
			result = result && rightResult
			result = result && (rightMi > node.Val)
			ma = max(ma, rightMa)
		}

		return mi, ma, result
	}

	if root == nil {
		return true
	}

	_, _, result := check(root)

	return result
}
