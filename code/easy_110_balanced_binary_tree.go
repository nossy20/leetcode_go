/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */

package code

// @lc code=start
//lint:ignore U1000 for leetcode
func isBalanced(root *TreeNode) bool {
	// edge case
	if root == nil {
		return true
	}

	// recursive
	res := depth(root.Left) - depth(root.Right)
	// 次以降の枝もバランスされているかをチェックする
	return (res <= 1 && res >= -1) && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// recursive
	// maxはmaximun_depth_of_binary_tree.goにある
	return max110(depth(root.Left), depth(root.Right)) + 1
}

func max110(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
