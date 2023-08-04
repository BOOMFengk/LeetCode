/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	up,down:=1,1
	for i := 1; i < len(nums); i++ {
		cur:=nums[i]-nums[i-1]
		if cur>0{
			up=down+1
		}else if cur<0{
			down = up + 1
		}
		
	}
	return max(up,down)
}

func max(i,j int)int{
	if i>j{
		return i
	}
	return j
}

// @lc code=end

