/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package code

// @lc code=start
//lint:ignore U1000 It's ok to have unused code
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// @lc code=end
