package greedy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func minCameraCover(root *TreeNode) int {
	res = 0
	if postOrderTraverse(root) == 0 {
		res++
	}
	return res
}

func postOrderTraverse(root *TreeNode) int {
	if root == nil {
		return 2
	}
	left := postOrderTraverse(root.Left)
	right := postOrderTraverse(root.Right)
	if left == 0 || right == 0 {
		res++
		return 1
	}
	if left == 2 && right == 2 {
		return 0
	}
	if left == 1 || right == 1 {
		return 2
	}
	return 0
}
