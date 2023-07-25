/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	result:=make([]int, 0)
	mp := map[int]int{}

	var trace func (root *TreeNode)  
	trace = func (root *TreeNode)  {
		if root == nil{
			return
		}
		trace(root.Left)
		mp[root.Val]++
		trace(root.Right)
		
	}	
	trace(root)

	cnt:=0
	for val, sum := range mp {
		if sum>cnt{
			cnt = sum
			result = []int{val}
		}else if sum == cnt{
			result = append(result,val)
		}
		
	}
	return result

}
// @lc code=end

