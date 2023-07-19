/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree root.
 * type Treeroot struct {
 *     Val int
 *     Left *Treeroot
 *     Right *Treeroot
 * }
 */
func inorderTraversal(root *TreeNode) (res []int) {
	var Traserval func(root *TreeNode)
	Traserval = func (root *TreeNode)  {
		if root != nil{
			Traserval(root.Left)
			res = append(res,root.Val)
			Traserval(root.Right)
		}
	}
	Traserval(root)
	return res

}
// @lc code=end

