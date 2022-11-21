/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

package code

// @lc code=start
//lint:ignore U1000 func addBinary(a string, b string) string {
func addBinary(a string, b string) string {
	result := []byte{}
	sum := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || sum != 0; i, j = i-1, j-1 {
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		result = append(result, byte(sum)%2+'0')
		// 繰り上げ
		// i < 0, j < 0の時はsumは0に収束する(繰り上げ分を処理して終了)
		sum /= 2
	}
	// reverse
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

// @lc code=end
