/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func intersection(nums1 []int, nums2 []int) []int {
	// 一時的なmapを作成する
	t := make(map[int]bool)
	// num1の要素をmapに追加する
	for _, v := range nums1 {
		t[v] = true
	}
	// num2の要素でmapに存在するものを抽出する
	r := make([]int, 0)
	for _, v := range nums2 {
		if t[v] {
			r = append(r, v)
			t[v] = false
		}
	}
	return r
}

// @lc code=end
