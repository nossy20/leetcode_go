/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func isUgly(n int) bool {
	// edge case
	if n <= 0 {
		return false
	}

	// 2で割り切れる場合は、2で割る
	for n%2 == 0 {
		n /= 2
	}

	// 3で割り切れる場合は、3で割る
	for n%3 == 0 {
		n /= 3
	}

	// 5で割り切れる場合は、5で割る
	for n%5 == 0 {
		n /= 5
	}

	// すべて割り切れていれば1になる
	if n == 1 {
		return true
	}
	return false
}

// @lc code=end
