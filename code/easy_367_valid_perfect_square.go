/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func isPerfectSquare(num int) bool {
	// edge case
	if num == 1 {
		return true
	}

	// 整数のみが対象なので、+1でインクリメントしつつ、平方根であるかを判定する
	for i := 2; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}

// @lc code=end
