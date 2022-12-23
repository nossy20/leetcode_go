/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

package code

// @lc code=start
//lint:ignore U1000 leetcode
func containsDuplicate(nums []int) bool {
	// 一時的な辞書を用意する
	t := map[int]int{}
	for _, v := range nums {
		// すでに辞書に存在する場合は重複があるのでtrueを返す
		if _, ok := t[v]; ok {
			return true
		}
		// 辞書に存在しない場合は追加する。バリューは何でも良い
		t[v] = 1
	}
	return false
}

// @lc code=end
