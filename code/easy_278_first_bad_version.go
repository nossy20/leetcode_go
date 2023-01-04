/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

package code

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

//lint:ignore U1000 LeetCode
func firstBadVersion(n int) int {
	// 二分探索
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end
