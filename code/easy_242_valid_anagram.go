/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func isAnagram(s string, t string) bool {
	// s, tの文字数を辞書でカウント
	a := make(map[rune]int)
	b := make(map[rune]int)
	for _, c := range s {
		if _, ok := a[c]; !ok {
			a[c] = 1
			continue
		}
		a[c] += 1
	}
	for _, c := range t {
		if _, ok := b[c]; !ok {
			b[c] = 1
			continue
		}
		b[c] += 1
	}

	// s, tの文字数が一致するかチェック
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// @lc code=end
