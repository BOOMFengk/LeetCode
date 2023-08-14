/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
//字符串不能修改
func monotoneIncreasingDigits(n int) int {
	temp:=strconv.Itoa(n)
	str:=[]byte(temp)
	flag:=len(str)
	for i := len(str)-1; i >0; i-- {
		if str[i-1]>str[i]{
			str[i-1]--
			flag=i
		}
		
	}
	for i := flag; i < len(str); i++ {
		str[i]='9'
	}
	temp=string(str)
	ans,_:=strconv.Atoi(temp)
	return ans

}
// @lc code=end

