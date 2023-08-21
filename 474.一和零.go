package main
/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp:= make([][]int, m+1)
	for i, _ := range dp {
		dp[i]-make([]int, n+1)
	}

	for i := 0; i < len(str); i++ {
		zeroNum,oneNum :=0,0
		zeroNum = strings.Count(strs[i],"0")
		oneNum = len(strs[i])-zeroNum
	}

	for j := m; j >=zeroNum; j-- {
		for k := n; k > oneNum; k-- {
			dp[j][k]=max(dp[j][k],dp[j-zeroNum][k-oneNum]+1)
		}
			
	}
	return dp[m][n]
}
func max(a,b int)int {
	if a>b{
		return a
	}
	return b
	
}
// @lc code=end

