/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums)==1{
		return 0
	}
	next,ans,cur:=0,0,0
	for i := 0; i < len(nums); i++ {
		next=max(next,nums[i]+i)
		if cur==i{
			ans++
			cur = next
			if next >=len(nums)-1{
				break
			}
		}
	}
	return ans

}

func max(i,j int)int{
	if i>j{
		return i 
	}
	return j
}
// @lc code=end

