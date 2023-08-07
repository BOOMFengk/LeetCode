/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	cnt:=0
	for i := 0; i < len(nums); i++ {
		if i>cnt{
			break
		}
		cnt = max(cnt,i+nums[i])
		if cnt>=len(nums)-1{
			return true
		}
		
	}
	return false

}

func max(i,j int)int{
	if i>j{
		return i
	}
	return j 
}
// @lc code=end

