/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) (ans [][]string ){
	var trace func(cur []string,idx int)
	trace =func(cur []string,idx int){
		if idx == len(s){
			temp:=make([]string, len(cur))
			copy(temp,cur)
			ans = append(ans,temp)
		}
		for i := idx; i < len(s); i++ {
			if check(idx,i,s){
				next:=append(cur,string(s[idx:i+1]))
				trace(next,i+1)
			}
			
		}
	}
	trace([]string{},0)
	return



}
func check(i,j int,s string)bool  {
	for i<=j{
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true
	
}
// @lc code=end

