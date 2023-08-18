/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	target:= 0
	for _, v := range nums {
		target+=v
	}
	if target%2!=0{
		return false
	}
	target/=2

	dp:=make([]int, target+1)
	
	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j]<dp[j-num]+num{
				dp[j]=dp[j-num]+num
			}	
		}
	}
	return dp[target]==target


}
// @lc code=end

