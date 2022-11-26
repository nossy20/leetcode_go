/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */

// reference: https://blog.devgenius.io/the-main-idea-of-implementing-a-binary-tree-with-golang-556fac53ced4

package code

// @lc code=start

//lint:ignore U1000 func inorderTraversal(root *TreeNode) []int {
func inorderTraversal(root *TreeNode) []int {
	out := []int{}
	if root != nil {
		out = append(out, inorderTraversal(root.Left)...)
		out = append(out, root.Val)
		out = append(out, inorderTraversal(root.Right)...)
	}
	return out
}

// @lc code=end
