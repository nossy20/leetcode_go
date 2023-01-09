/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func reverseVowels(s string) string {
	r := []rune{'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O'}
	o := []rune(s)
	for i, j := 0, len(s)-1; i < j; {
		// iが母音でない時はiをインクリメント
		if !isVowel(o[i], r) {
			i++
		}

		// jが母音でない時はjをデクリメント
		if !isVowel(o[j], r) {
			j--
		}

		// 両方が母音の時は入れ替えてiとjをインクリメント・デクリメント
		if isVowel(o[i], r) && isVowel(o[j], r) {
			o[i], o[j] = o[j], o[i]
			i++
			j--
		}
	}
	return string(o)
}

func isVowel(c rune, r []rune) bool {
	for _, v := range r {
		if c == v {
			return true
		}
	}
	return false
}

// @lc code=end
