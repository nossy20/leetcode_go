/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

package code

// @lc code=start
//lint:ignore U1000 func convertToTitle(columnNumber int) string {
func convertToTitle(columnNumber int) string {
	// 結果を格納する配列
	r := []byte{}
	for columnNumber > 0 {
		// 0-25の範囲で、A-Zを表すため-1する
		columnNumber--
		// 26進数のように、26で割った余りを配列の先頭に追加する
		r = append([]byte{byte('A' + columnNumber%26)}, r...)
		// 上の位の桁を計算するため、26で割る
		columnNumber /= 26
	}
	return string(r)
}

// @lc code=end
