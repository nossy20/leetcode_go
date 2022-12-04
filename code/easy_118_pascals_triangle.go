/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

package code

// @lc code=start
//lint:ignore U1000 ignore unused code
func generate(numRows int) [][]int {
	// edge cases
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	// generate and init
	out := make([][]int, numRows)
	// init
	out[0] = []int{1}
	out[1] = []int{1, 1}
	for i := 3; i <= numRows; i++ {
		list := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				list[j] = 1
			} else {
				list[j] = out[i-2][j-1] + out[i-2][j]
			}
		}
		out[i-1] = list
	}
	return out
}

// @lc code=end
