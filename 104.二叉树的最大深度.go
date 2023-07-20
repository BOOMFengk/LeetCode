/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	if root ==nil{
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	hight :=max(left,right)+1 
	return hight
}
func max(i,j int)  int{
	if i>j{
		return i
	}
	return j
}

// @lc code=end

