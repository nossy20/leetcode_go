/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func moveZeroes(nums []int) {
	// iは0以外の数を探すためのインデックス
	// jは0を探すためのインデックス
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return
}

// @lc code=end
