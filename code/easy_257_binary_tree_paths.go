/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */

package code

import "strconv"

// @lc code=start
//lint:ignore U1000 LeetCode
func binaryTreePaths(root *TreeNode) []string {
	// edge case
	if root == nil {
		return []string{}
	}

	// 二分木のルートから葉ノードまでのパスを作成
	paths := []string{}
	createPath(root, "", &paths)

	return paths
}

func createPath(root *TreeNode, path string, paths *[]string) {
	if root == nil {
		return
	}
	// 終端する
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path+strconv.Itoa(root.Val))
		return
	}
	// 左右の子ノードを再帰的に探索
	createPath(root.Left, path+strconv.Itoa(root.Val)+"->", paths)
	createPath(root.Right, path+strconv.Itoa(root.Val)+"->", paths)
}

// @lc code=end
