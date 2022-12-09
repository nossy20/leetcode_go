/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

package code

import "fmt"

// @lc code=start
//lint:ignore U1000 function is unused
func singleNumber(nums []int) int {
	r := 0
	for _, n := range nums {
		// XOR排他的論理和を取る
		r ^= n
		fmt.Printf("Result: %08b\n", r)
	}
	return r
}

// @lc code=end
