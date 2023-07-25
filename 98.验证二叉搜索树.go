/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode

	var trace func (root *TreeNode)bool  
		trace = func (root *TreeNode)bool  {
			if root ==nil {
		  	return true
			}
			left := trace(root.Left)
			if prev !=nil && prev.Val >=root.Val{
				return false
			}	else{
				prev = root
			}
			right := trace(root.Right)
			

			return left && right
		}

		
		
	
	return trace(root)
}
// @lc code=end

