

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	res:=make([]int, 0)
	var recursion func(root *TreeNode)  
	recursion = func(root *TreeNode){
		if root == nil{
			return
		}
		res = append(res,root.Val)
		recursion(root.Left)
		recursion(root.Right)
	}
	recursion(root)
	return res

}
// @lc code=end

