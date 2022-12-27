/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

package code

import "strconv"

// @lc code=start
//lint:ignore U1000 leetcode
func summaryRanges(nums []int) []string {
	// edge case
	if len(nums) == 0 {
		return []string{}
	}

	res := []string{}
	s := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		} else {
			if s == nums[i-1] {
				res = append(res, strconv.Itoa(s))
			} else {
				res = append(res, strconv.Itoa(s)+"->"+strconv.Itoa(nums[i-1]))
			}
			s = nums[i]
		}
	}
	// 最後の要素の判定
	if s == nums[len(nums)-1] {
		res = append(res, strconv.Itoa(s))
	} else {
		res = append(res, strconv.Itoa(s)+"->"+strconv.Itoa(nums[len(nums)-1]))
	}
	return res
}

// @lc code=end
