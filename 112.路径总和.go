/*
 * @lc app=leetcode.cn id=112 lang=golang
 * [112] 路径总和
 */
 
 // @lc code=start
 /**
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	var trace func (*TreeNode,int)bool
	trace  = func (root *TreeNode,cnt int)bool  {
		if root ==nil{
			return false
		}
		cnt += root.Val
		if root.Left ==nil && root.Right ==nil &&cnt ==targetSum{
			return true
		}
		return trace(root.Left,cnt)||trace(root.Right,cnt)
	}
		return trace(root,0)
}
// @lc code=end

