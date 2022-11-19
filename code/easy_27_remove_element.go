/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

package code

// @lc code=start
//lint:ignore U1000 func removeElement(nums []int, val int) int {
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

// @lc code=end
