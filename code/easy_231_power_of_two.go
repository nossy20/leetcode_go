/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func isPowerOfTwo(n int) bool {
	// edge case
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true // 2^0
	}
	for n > 1 {
		// あまりが0でない場合は2の累乗ではない
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}

// @lc code=end
