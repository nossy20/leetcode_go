/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

package code

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

//lint:ignore U1000 LeetCode
func guessNumber(n int) int {
	low := 0
	high := n
	cur := (low + high) / 2

	// 最初から等しい場合
	if guess(cur) == 0 {
		return cur
	}

	// 二分探索
	for low <= high {
		if guess(cur) == 0 {
			return cur
		} else if guess(cur) == -1 {
			high = cur - 1
		} else {
			low = cur + 1
		}
		cur = (low + high) / 2
	}
	return 0 // ここには来ない
}

// @lc code=end
