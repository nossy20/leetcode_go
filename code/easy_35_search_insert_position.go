/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

package code

// @lc code=start
//lint:ignore U1000 func searchInsert(nums []int, target int) int {
func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return len(nums)
}

// @lc code=end
