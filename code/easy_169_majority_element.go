/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

package code

// @lc code=start
//lint:ignore U1000 LeetCode
func majorityElement(nums []int) int {
	cnt_map := map[int]int{}
	for _, num := range nums {
		cnt_map[num]++
	}
	for key, val := range cnt_map {
		if val > len(nums)/2 {
			return key
		}
	}
	return 0
}

// @lc code=end
