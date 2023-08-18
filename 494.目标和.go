/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	//A-B = target
	//A+B = SUM
	//A+B+A-B =target +sum
	//2A = target +sum
	//A = (target+sum)/2
	sum:=0
	for _, v := range nums {
		sum+=v
	}
	n:=(target+sum)/2
	
	if abs(target)>sum ||(target+sum)%2==1{
		return 0
	}

	dp:=make([]int, n+1)
	dp[0]=1
	for i := 0; i < len(nums); i++ {
		for j := n; j >= nums[i]; j-- {
			dp[j]=dp[j]+dp[j-nums[i]]	
		}
	}
	return dp[n]
}
func abs(i int) int {
	if i<0{
		return -i
	}
	return i
	
}
// @lc code=end

