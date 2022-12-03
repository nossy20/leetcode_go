/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 */

package code

// @lc code=start
//lint:ignore U1000 for leetcode
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// recursive
	return near(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 最小の距離を返却する
func near(a, b int) int {
	// 両方が末端の場合は、0が返却される
	if a == 0 && b == 0 {
		return 0
	}
	// 一方が深さを保つ場合は、深さを返却する
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	// 両方が深さを保つ場合は、小さい方を返却する
	return min(a, b)
}

// 小さい方を返却する
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
