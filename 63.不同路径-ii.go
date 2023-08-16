/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	//写了几次二维数组的代码  make好像只能初始化一个长度，另一个长度
	//需要dp[]=make([]int,n+1)这种.
	ob :=obstacleGrid
	m,n:=len(ob),len(ob[0])
	dp:=make([][]int, m)
	for i := range dp {
		dp[i]=make([]int,n)
	}
	for i := 0; i < m; i++ {
		if ob[i][0]!=1{
			dp[i][0] = 1
		}else {
			break
		}
	}
	for j := 0; j < n; j++ {
		if ob[0][j]!=1{
			dp[0][j]=1
		}else{
			break
		}
		
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if ob[i][j]!=1{
				dp[i][j]=dp[i-1][j]+dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]

}
// @lc code=end

