/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

package code

// @lc code=start
//lint:ignore U1000 func isValid(s string) bool {
func isValid(s string) bool {
	// 0または奇数の場合はfalse
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	// 1文字目が閉じ括弧の場合はfalse
	if s[0] == ')' || s[0] == '}' || s[0] == ']' {
		return false
	}

	// 開き括弧と閉じ括弧の対応を定義
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 開き括弧を格納するスタック
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		// 開き括弧の場合はスタックに追加
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			// 閉じ括弧の場合はスタックの最後の要素と対応する開き括弧か確認
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			// 対応する開き括弧の場合はスタックから削除
			stack = stack[:len(stack)-1]
		}
	}

	// スタックが空の場合はtrue
	return len(stack) == 0
}

// @lc code=end
