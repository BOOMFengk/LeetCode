
/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(s string) string {
	stack := make([]byte,0)
	for i := range s {
		if len(stack)==0 || stack[len(stack)-1]!=s[i]{
			stack = append(stack,s[i])
		}else {
			stack = stack[:len(stack)-1]
		}
		
	}
	return string(stack)
}
// @lc code=end

