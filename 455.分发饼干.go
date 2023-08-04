/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans:=0
	i:=len(s)-1
	j:=len(g)-1
	for ; i >=0; i-- {
		have:=s[i]
		for ; j >=0; j-- {
			need:=g[j]
			if have>=need{
				ans++
				j--
				break
			}	
		}
	}
	return ans
}
// @lc code=end

