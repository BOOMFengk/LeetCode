/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) (ans[][]int) {
	mp := map[int]struct{}{}
	var trace func(cur []int)
	trace = func(cur []int)  {
		if len(cur)==len(nums){
			temp:=make([]int, len(cur))
			copy(temp,cur)
			ans=append(ans,temp)
		}
		mmp:=map[int]struct{}{}
		for idx, v := range nums {
			if _,ok:=mp[idx];ok {
				continue
			}
			if _,ok:=mmp[v];ok{
				continue
			}
			mmp[v]=struct{}{}
			mp[idx]=struct{}{}
			next:=append(cur,v)
			trace(next)
			delete(mp,idx)
		}

	}
	trace([]int{})
	return

}
// @lc code=end

