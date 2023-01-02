/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func missingNumber(nums []int) int {
	// 0 + 1 + 2 + ... + n = n * (n + 1) / 2
	// すべての足し算は、n * (n + 1) / 2 で表せる
	sum := len(nums) * (len(nums) + 1) / 2
	total := 0
	for _, n := range nums {
		total += n
	}
	return sum - total
}

// @lc code=end
