/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	if root ==nil{
		return 0
	}
	left,right := root.Left,root.Right
	leftdepth,rightdepth :=0,0
	for left!=nil{
		left =left.Left
		leftdepth++
	}
	for right !=nil{
		right = right.Right
		rightdepth++
	}

	if leftdepth==rightdepth{
		return 2<<leftdepth - 1
	}
	return countNodes(root.Left)+countNodes(root.Right)+1
}
// @lc code=end

