/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) (ans[][]int) {
	var trace func (cur []int ,idx int)
	trace = func (cur []int,idx int)  {
		if true{
			temp:=make([]int, len(cur))
			copy(temp,cur)
			ans = append(ans,temp)
		}
		for i := idx; i < len(nums); i++ {
			next := append(cur,nums[i])
			trace(next,i+1)
			
		}
	}  
	trace([]int{},0)
	return

}
// @lc code=end

