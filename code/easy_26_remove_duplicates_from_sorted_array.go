/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

package code

// @lc code=start
//lint:ignore U1000 func removeDuplicates(nums []int) int {
func removeDuplicates(nums []int) int {
	o := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			nums[o] = nums[i]
			o++
		}
	}
	return o
}

// @lc code=end
