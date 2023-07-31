/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) (ans [][]int ){
	sort.Ints(candidates)
	var trace func(cur []int,sum int, idx int)
	trace = func (cur []int,sum int,idx int){
		if sum >target{
			return
		}
		if sum == target{
			temp:=make([]int, len(cur))
			copy(temp,cur)
			ans = append(ans,temp)
		}
		mp:=map[int]bool{}
		for i := idx; i < len(candidates); i++ {
			if _,ok:=mp[candidates[i]];ok{
				continue
			}
			mp[candidates[i]]=true
			next:=append(cur,candidates[i])
			trace(next,sum+candidates[i],i+1)
		}
	}  
	trace([]int{},0,0)
	return

}
// @lc code=end

