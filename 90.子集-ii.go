/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) (ans[][]int ){
	sort.Ints(nums)
	var trace func (cur []int ,idx int)
	trace = func (cur []int,idx int)  {
		if true{
			temp:=make([]int, len(cur))
			copy(temp,cur)
			ans = append(ans,temp)
		}
		mp:=map[int]struct{}{}
		for i := idx; i < len(nums); i++ {
			if _,ok:=mp[nums[i]];ok{
				continue
			}
			mp[nums[i]]=struct{}{}
			next := append(cur,nums[i])
			trace(next,i+1)
			
		}
	}  
	trace([]int{},0)
	return

}
// @lc code=end

