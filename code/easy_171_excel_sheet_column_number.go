/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func titleToNumber(columnTitle string) int {
	// 結果を格納する配列
	// バイトに変換する "AB" = []byte{65, 66}
	bs := []byte(columnTitle)
	c_num := 0
	for i := 0; i < len(bs); i++ {
		// A-Z => 1-26で計算する
		t := int(bs[i]-'A') + 1
		// 桁がシフトするため、26をかける
		// AA => 1 * 26 + 1
		// AAA => 27(AA) * 26 + 1
		c_num = 26*c_num + t
	}
	return c_num
}

// @lc code=end
