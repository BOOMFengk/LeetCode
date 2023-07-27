/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}
	midIdx := len(nums)/2
	root := &TreeNode{nums[midIdx],nil,nil}
	root.Left = sortedArrayToBST(nums[:midIdx])
	root.Right = sortedArrayToBST(nums[midIdx+1:])
	return root
}
// @lc code=end

