/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start

//这题跟518差不多  518是组合数，这个是排列数
func combinationSum4(nums []int, target int) int {
	dp:=make([]int, 5001)
	n:=len(nums)
	dp[0]=1
	for j := 0; j <= target; j++ {
		for i := 0; i < n; i++ {
			if j>=nums[i]{
				dp[j]+=dp[j-nums[i]]
			}
			
		}
		
	}
	return dp[target]

}
// @lc code=end

