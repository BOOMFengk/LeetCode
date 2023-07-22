/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
	var trace func(*TreeNode ,bool)int
		trace = func (root *TreeNode,isL bool)int  {
			if root ==nil{
				return 0
			}
			if root.Left ==nil && root.Right ==nil&& isL{
				return root.Val
			}
			left:=trace(root.Left,true)
			right :=trace(root.Right,false)
			return left + right
		}
	
	return trace(root,false)

}
// @lc code=end

