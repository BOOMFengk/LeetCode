/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) (ans[][]int ){
	var trace func(cur []int,idx int)
	trace = func(cur []int,idx int)  {
		if len(cur)>=2{
			temp:=make([]int, len(cur))
			copy(temp,cur)
			ans=append(ans,temp)
		}
		mp :=map[int]struct{}{}
		for i := idx; i <len(nums) ; i++ {
			if _,ok:=mp[nums[i]];ok{
				continue
			}
			mp[nums[i]]=struct{}{}
			if len(cur)==0||cur[len(cur)-1]<=nums[i]{
				next:=append(cur,nums[i])
				trace(next,i+1)
			}
		}
		
	}
	trace([]int{},0)
	return

}
// @lc code=end

