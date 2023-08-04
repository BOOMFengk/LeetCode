/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// maxn,sum:=-0x3f3f3f3f, 0
	// for i := 0; i < len(nums); i++ {
	// 	sum+=nums[i]
	// 	if sum>maxn{
	// 		maxn =sum
	// 	}
	// 	if sum<0{
	// 		sum=0
	// 	}
		n:=len(nums)
		dp:=make([]int, n+1)
		dp[0]=0
		ans:=math.MinInt32
		for i := 1; i <=n; i++ {
			dp[i]=max(nums[i-1],dp[i-1]+nums[i-1])
			ans= max(ans,dp[i])
		}
		return ans
		
	}

	func max(i,j int)int  {
		if i>j{
			return i
		}
		return j
		
	}

	

// @lc code=end

