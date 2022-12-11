/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */

package code

// @lc code=start
//lint:ignore U1000 leetcode
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// recursive
	return append([]int{root.Val}, append(preorderTraversal(root.Left), preorderTraversal(root.Right)...)...)
}

// @lc code=end
