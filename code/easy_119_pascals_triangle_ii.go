/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

package code

// @lc code=start
//lint:ignore U1000 ignore unused code
func getRow(rowIndex int) []int {
	// edge cases
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	// generate and init
	out := make([][]int, rowIndex+1)
	// init
	out[0] = []int{1}
	out[1] = []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		list := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				list[j] = 1
			} else {
				list[j] = out[i-1][j-1] + out[i-1][j]
			}
		}
		out[i] = list
	}
	return out[rowIndex]
}

// @lc code=end
