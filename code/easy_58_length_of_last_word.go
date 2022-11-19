/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

package code

// @lc code=start
//lint:ignore U1000 func lengthOfLastWord(s string) int
func lengthOfLastWord(s string) int {
	// cnt := 文字数をカウントする
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			// 文字列をカウントする
			cnt++
		} else if cnt > 0 {
			// 空文字かつ、文字列が存在した場合はカウントを返す
			return cnt
		}
		// 空文字かつ、文字列が存在しない場合は何もしない
		continue
	}
	return cnt
}

// @lc code=end
