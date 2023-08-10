/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) (result [][]int) {
	sort.Slice(intervals,func (i,j int) bool {
		return intervals[i][0]<intervals[j][0]
	})
	left:=intervals[0][0]
	right:=intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		curLeft:=intervals[i][0]
		
		if right>=curLeft {
			right=max(right,intervals[i][1])
		}else{
			result=append(result,[]int{left,right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	result = append(result,[]int{left,right})
	return
}

func max(i,j int) int {
	if i>j{
		return i
	}
	return j
}
// @lc code=end

