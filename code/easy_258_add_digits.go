/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

package code

import "strconv"

// @lc code=start
//lint:ignore U1000 LeetCode
func addDigits(num int) int {
	// 一桁だった場合はそのまま返す
	if len(strconv.Itoa(num)) == 1 {
		return num
	}

	// 10以上の場合は、各桁の和を求める
	return addDigits(sumDigits(num))
}

func sumDigits(num int) int {
	// 各桁の和を求める
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// @lc code=end
