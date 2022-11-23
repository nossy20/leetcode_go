/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

package code

import "fmt"

// @lc code=start
//lint:ignore U1000 func climbStairs(n int) int
func climbStairs(n int) int {
	// dp 動的計画法を用いたアプローチ
	// 段をのぼる方法は1段上がるか2段上がるかの2通りしかない
	// n段目に達する方法は、n-1段目から1段上がるか、n-2段目から2段上がるか
	// したがって、n段目に達する方法はn-1段目に達する方法の総和+n-2段目に達する方法の総和
	// 初期化, 2段目までは初期として扱う
	// n = 1 -> return 1;
	// n = 2 -> return 2;
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 0からnまでのn+1サイズのテーブルを用意
	dp := make([]int, n+1)
	// 初期化
	dp[1] = 1
	dp[2] = 2
	// 3段以降の処理
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		fmt.Println(dp)
	}

	return dp[n]

}

// @lc code=end
