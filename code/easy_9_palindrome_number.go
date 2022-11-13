/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

package code

import (
	"strconv"
)

// @lc code=start
//lint:ignore U1000 ignore unused code
func isPalindrome(x int) bool {
	f := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	s := strconv.Itoa(x)
	return s == f(s)
}

// @lc code=end
