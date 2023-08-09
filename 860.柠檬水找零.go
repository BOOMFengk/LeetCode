/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	ten,five:=0,0
	for i := 0; i < len(bills); i++ {
		if bills[i]==5 {
			five++
		}else if bills[i]==10 {
			if five==0 {
				return false	
			}
			ten++
			five--
		}else{
			if ten>=1&&five>=1{
				ten--
				five--
			}else if five>=3{
				five-=3
			}else{
				return false
			}
		}
	}
	return true

}
// @lc code=end

