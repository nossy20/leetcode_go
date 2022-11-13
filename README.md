# 環境構築

- [LeetCode を VS Code と Go で解くための環境づくり](https://zenn.dev/sadah/articles/ede353e39c0500)
- [vscode-leetcode で快適 LeetCode 生活](https://zenn.dev/ryo_kawamata/articles/introduce-vscode-leetcode)

# 問題を解く前に

1. package を宣言
2. unused 警告を無視

```golang
/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package code
// @lc code=start
//lint:ignore U1000 It's ok to have unused code
func twoSum(nums []int, target int) []int {
    {{ your code }}
}
```
