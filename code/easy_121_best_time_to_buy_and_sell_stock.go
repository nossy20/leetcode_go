/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

package code

// @lc code=start
//lint:ignore U1000 ignore unused code
func maxProfit(prices []int) int {
	max := 0     // from constraints
	min := 10000 // from constraints
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}

// @lc code=end
