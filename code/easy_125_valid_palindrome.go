/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

package code

import "unicode"

// @lc code=start
//lint:ignore U1000 func isPalindrome(s string) bool {
func isPalindrome(s string) bool {
	// 与えられた文字列を前後から順に比較していく
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// 比較文字列が、アルファベットか数字でない場合は、スキップする(前方)
		for i < j && !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
		}
		// 比較文字列が、アルファベットか数字でない場合は、スキップする(後方)
		for i < j && !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
		}
		// 文字列を比較して異なる場合はその時点でfalseを返却して終了
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
	}
	return true
}

// @lc code=end
