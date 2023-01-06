/*
 * @lc app=leetcode id=292 lang=golang
 *
 * [292] Nim Game
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func canWinNim(n int) bool {
	// 4で割り切れるときは必ず負ける
	// 4で割り切れないときは必ず勝てる
	if n%4 == 0 {
		return false
	}
	return true
}

// @lc code=end
