/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start

var(
	path []int
	res [][]int
)
func combine(n int, k int)[][]int {
	path,res = make([]int, 0,k),make([][]int,0)
	var trace func(n int ,start int)

	trace = func (n int ,start int)  {
	if len(path)==k{
		tmp:=make([]int, k)
		copy(tmp,path)
		res =append(res,tmp)
		return
	}
	for i := start; i <=n; i++ {
		// if n-i+1 <k-len(path){
		// 	break
		// }
		path = append(path,i)
		trace(n,i+1)
		path=path[:len(path)-1]
	}
	
}
	trace(n,1)
	return res
}



// @lc code=end

