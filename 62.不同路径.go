/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	//写了几次二维数组的代码  make好像只能初始化一个长度，另一个长度
	//需要dp[]=make([]int,n+1)这种.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range dp {
		dp[i][0] = 1

	}
	for i := range dp[0] {
		dp[0][i] = 1

	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]

		}
	}
	return dp[m-1][n-1]

}

// @lc code=end

