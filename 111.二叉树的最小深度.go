/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root ==nil{
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if root.Left ==nil && root.Right !=nil{
		return right+1
	}else if root.Left !=nil &&root.Right ==nil{
		return left+1
	}
	return min(left,right)+1

}

func min(i,j int) int {
	if i>j{
		return j
	}
	return i
	
}
// @lc code=end

