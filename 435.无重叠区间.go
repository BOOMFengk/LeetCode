/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	
	sort.Slice(intervals,func (i,j int) bool {
		return intervals[i][1]<intervals[j][1]
	})
	res:=1
	oldr:=intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		newr:=intervals[i][0]
		if oldr<=newr {
			res++
			oldr=intervals[i][1]
		}
		
	}
	return len(intervals)-res
}
// @lc code=end

