/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) (ans [][]int) {
	var trace func(sum int,cur []int,idx int)
	trace = func(sum int,cur []int,idx int){
		if sum > target{
			return
		}if sum==target {
			temp:=make(int[], len(cur))
			copy(temp,cur)
			ans = append(ans,temp)
			
		}
		for i := idx; i < len(candidates); i++ {
			next:=append(cur,candidates[i])
			trace(sum+candidates[i],next,i)
			
		}
	}
	trace(0,{}int{},0)
	return
}
// @lc code=end

