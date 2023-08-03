/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) (ans [][]int) {
	mp := map[int]struct{}{}
	var trace func(cur []int)
	trace = func(cur []int)  {
		if len(cur)==len(nums){
			temp:=make([]int, len(cur))
			copy(temp,cur)
			ans=append(ans,temp)
		}
		for idx, v := range nums {
			if _,ok:=mp[idx];ok {
				continue
			}
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

