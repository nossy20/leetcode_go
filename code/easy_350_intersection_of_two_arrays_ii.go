/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func intersect(nums1 []int, nums2 []int) []int {
	// 一時的にnums1の要素を格納するmapを作成
	t := make(map[int]int)
	for _, v := range nums1 {
		t[v]++
	}
	// nums2の要素をmapから探し、あればrに追加
	r := []int{}
	for _, v := range nums2 {
		if t[v] > 0 {
			r = append(r, v)
			t[v]--
		}
	}
	return r
}

// @lc code=end
