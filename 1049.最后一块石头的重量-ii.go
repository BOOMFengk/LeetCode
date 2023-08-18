/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	target:=0
	for _, v := range stones {
		target+=v
	}
	flag:=0
	if target%2!=0{
		flag=1
	}
	sum:=target
	target/=2
	dp:=make([]int, target+1)
	n:=len(stones)
	for i := 0; i <n ; i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j]=max(dp[j],dp[j-stones[i]]+stones[i])	
		}
	}
		if dp[target]==target&&flag==0{
			return 0
	}
	return sum - 2*dp[target]

}
func max(i,j int) int {
	if i>j{
		return i
	}
	return j
	
}
// @lc code=end

