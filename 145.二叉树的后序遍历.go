/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) (res []int) {
	var Traserval func(root *TreeNode)
	Traserval = func (root *TreeNode)  {
		if root != nil{
			Traserval(root.Left)
			
			Traserval(root.Right)
			res = append(res,root.Val)
		}
	}
	Traserval(root)
	return res

}
// @lc code=end

