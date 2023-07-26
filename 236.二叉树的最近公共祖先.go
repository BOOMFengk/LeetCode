/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	if root == p || root == q{
		return root
	}
	left , right := lowestCommonAncestor(root.Left,p,q) ,lowestCommonAncestor(root.Right,p,q)
	if left!=nil &&right!=nil{
		return root
	}

	if left ==nil && right !=nil{
		return right
	}

	if left !=nil && right ==nil{
		return left
	}
	return nil
  
}
// @lc code=end

