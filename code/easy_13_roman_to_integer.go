/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

package code

// @lc code=start
//lint:ignore U1000 ignore unused code
func romanToInt(s string) int {
	// 対応するマップを作成する
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	// ループを回して足し算をする
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			// 次の文字列が大きい場合は自分の数だけ引けば良い -> IV = -1 + 5
			result -= romanMap[string(s[i])]
		} else {
			// 基本的には足し算
			result += romanMap[string(s[i])]
		}
	}
	return result
}

// @lc code=end
