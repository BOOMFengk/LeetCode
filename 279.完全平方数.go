/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	nums:=make([]int, 100)
	for i := range nums {
		nums[i]=(i+1)*(i+1)
	}
	dp:=make([]int, n+1)

	for i := range dp {
		dp[i]=math.MaxInt32
	}
	dp[0]=0

	for i := 0; i < 100; i++ {
		for j := nums[i]; j <= n; j++ {
			dp[j] = min(dp[j],dp[j-nums[i]]+1)	


		}

	}
	return dp[n]

}
func min(i,j int)int{
	if i<j{
		return i
	}
	return j
}
// @lc code=end

