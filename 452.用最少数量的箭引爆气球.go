/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	sort.Slice(points,func (i,j int)bool  {
		return points[i][0]<points[j][0]
	})
	res:=1
	for i := 1; i < len(points); i++ {
		old:=points[i-1][1]
		new:=points[i][0]

		if old<new{
			res++
		}else {
			points[i][1]=min(points[i][1],points[i-1][1])
		}
		
	}
	return res

}

func min(i,j int)int{
	if i<j{
		return i
	}
	return j
}
// @lc code=end

