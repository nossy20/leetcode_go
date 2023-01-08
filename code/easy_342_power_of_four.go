/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func isPowerOfFour(n int) bool {
	// 4のべき乗なので4で割り続けてあまりが存在せず、最終的に4になればOK
	for n > 0 && n%4 == 0 {
		n /= 4
	}
	return n == 1
}

// @lc code=end
