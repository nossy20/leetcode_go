/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

package code

// @lc code=start
//lint:ignore U1000 ignore unused code
func isPowerOfThree(n int) bool {
	// 3のべき乗なので3で割り続けてあまりが存在せず、最終的に1になればOK
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// @lc code=end
