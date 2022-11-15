/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

package code

// @lc code=start
//lint:ignore U1000 func longestCommonPrefix(strs []string) string {
func longestCommonPrefix(strs []string) string {
	prefix := ""
	// 配列が空の場合は空文字を返す
	if len(strs) == 0 {
		return prefix
	}

	// 垂直探索
	for i := 0; i < len(strs[0]); i++ {
		// 最初の文字列の1文字目を取得
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			// 配列の要素数がiを超えた場合は、空文字を返す
			if i == len(strs[j]) || strs[j][i] != c {
				return prefix
			}
		}
		prefix += string(c)
	}
	return prefix
}

// @lc code=end
