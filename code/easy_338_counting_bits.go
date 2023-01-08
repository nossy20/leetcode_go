/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func countBits(n int) []int {
	// 10進数を2進数に変換して、1の数を数える
	r := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cnt := 0
		for j := i; j > 0; j /= 2 {
			if j%2 == 1 {
				cnt++
			}
		}
		r[i] = cnt
	}
	return r
}

// @lc code=end
