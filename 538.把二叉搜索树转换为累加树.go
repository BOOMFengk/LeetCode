/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	var prev *TreeNode
	var trace func(*TreeNode)
	trace = func (root *TreeNode)  {
		if root ==nil{
			return 
		}
		trace(root.Right)
		if prev != nil{
			root.Val += prev.Val
		}
		prev = root
		trace(root.Left)
	}
	
	trace(root)
	return root

}
// @lc code=end

