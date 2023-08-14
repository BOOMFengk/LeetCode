/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	// if n < 2 {
	// 	return n
	// }
	// dp := make([]int, n+1)
	// dp[0] = 0
	// dp[1] = 1
	// for i := 2; i <= n; i++ {
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n]
	if n<=1{
		return n
	}
	return fib(n-1)+fib(n-2)
}

// @lc code=end

