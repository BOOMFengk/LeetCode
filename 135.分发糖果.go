/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	ans,n:=0,len(ratings)
	candys := make([]int, n)
	for i := range candys {
		candys[i]=1
		
	}
	for i := 1; i < n; i++ {
		if ratings[i]>ratings[i-1]{
			candys[i]=candys[i-1]+1
		}
		
	}
	for i := n-2; i >= 0; i-- {
		if ratings[i]>ratings[i] {
			candys[i]=max(candys[i],candys[i+1]+1)
		
		}
		
	}
	for _, v := range candys {
		ans+=v
		
	}
	return ans

}
func max(i,j int) int {

	if i>j {
		return i
	}
	return j
}
// @lc code=end

