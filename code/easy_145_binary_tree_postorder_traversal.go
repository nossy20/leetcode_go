/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */
package code

// @lc code=start
//lint:ignore U1000 for leetcode
func postorderTraversal(root *TreeNode) []int {
	if root != nil {
		return append(append(postorderTraversal(root.Left), postorderTraversal(root.Right)...), root.Val)
	}
	return nil
}

// @lc code=end
