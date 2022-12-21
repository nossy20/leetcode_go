/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

package code

// @lc code=start
//lint:ignore U1000 func isIsomorphic
func isIsomorphic(s string, t string) bool {
	// 2つのmapを作成して文字列がそれぞれ一対一に格納する
	a := map[string]string{}
	b := map[string]string{}

	for i := 0; i < len(s); i++ {
		if a[string(s[i])] == "" && b[string(t[i])] == "" {
			// aにもbにも文字列が格納されていない場合は1対1の対応をするように格納する
			a[string(s[i])] = string(t[i])
			b[string(t[i])] = string(s[i])
		} else if a[string(s[i])] != string(t[i]) || b[string(t[i])] != string(s[i]) {
			// どちらか一方にのみ格納されており、その対応が一致しない場合はfalseを返す
			return false

		}
	}
	return true
}

// @lc code=end
