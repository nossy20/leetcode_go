/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */

package code

// @lc code=start
//lint:ignore U1000 for leetcode
func sortedArrayToBST(nums []int) *TreeNode {
	// edge case
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}
	// find the middle index(sorted array)
	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
	return root
}

// @lc code=end
