/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

package code

// @lc code=start
//lint:ignore U1000 func plusOne(digits []int) []int
func plusOne(digits []int) []int {
	// 長さ
	n := len(digits)
	// 繰り上げして足す値
	carry := 0
	// 末尾から処理をしていく
	for i := n - 1; i >= 0; i-- {
		// 足し算の結果を格納する一時的な変数(d)
		var d int
		if i == n-1 {
			// 末尾の場合は1を足す
			d = digits[i] + carry + 1
		} else {
			// 末尾以外の場合は繰り上げを足す
			d = digits[i] + carry
		}

		// 足し算の結果を求める
		digits[i] = d % 10 // d<=9はそのまま値が入る。d=10の場合は0を入れる。末尾から書き直す。
		carry = d / 10     // d<=9の場合は0が入る。d=10の場合は1が入る。これが繰り上げとなる。
	}

	if carry == 0 {
		// 最後まで処理したときに、繰り上げが発生していない場合は、処理を終了する
		return digits
	}

	// 繰り上げが発生している場合は、先頭に1を追加する
	return append([]int{1}, digits...)
}

// @lc code=end
