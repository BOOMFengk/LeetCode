/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func (i,j int) bool {
		if people[i][0]==people[j][0]{
			return people[i][1]<people[j][1]
		}
		return people[i][0]>people[j][0]
	})

	ans:=make([][]int, len(people))
	for _, v := range people {
		k:=v[1]
		copy(ans[k+1:],ans[k:])
		ans[k]=v
		
	}
	return ans
	

}
// @lc code=end

