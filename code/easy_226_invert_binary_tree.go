/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 */

package code

// @lc code=start
//lint:ignore U1000 leetcode
func invertTree(root *TreeNode) *TreeNode {
	// edge case
	if root == nil {
		return nil
	}

	// 一時的に対比させる
	t_l := root.Left
	// ポインタを入れ替え
	root.Left = invertTree(root.Right)
	root.Right = invertTree(t_l)
	return root
}

// @lc code=end
