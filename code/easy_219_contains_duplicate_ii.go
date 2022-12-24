/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */
package code

// @lc code=start
//lint:ignore U1000 LeetCode
func containsNearbyDuplicate(nums []int, k int) bool {
	// 一時的な辞書を作成する
	t := map[int]int{}
	for i, v := range nums {
		// すでに存在する場合、そのインデックスとの差がk以下であればtrueを返す
		if n, ok := t[v]; ok && i-n <= k {
			return true
		}
		// なければ、辞書に追加する
		t[v] = i
	}
	return false
}

// @lc code=end
