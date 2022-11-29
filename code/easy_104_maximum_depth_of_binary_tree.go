/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */

package code

// @lc code=start
//lint:ignore U1000 for leetcode
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// recursive
	// 末端から上に向かって、深さを返却していく
	// 一番深い枝のみが返却される
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 多いきい方を返却する
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
