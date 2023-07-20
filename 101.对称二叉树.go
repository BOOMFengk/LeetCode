package LeetCode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	
	return compare(root.Left,root.Right)

}

func compare(left,right *TreeNode)bool{
	if left == nil && right !=nil{
		return false
	}else if left !=nil && right ==nil{
		return false
	}else if left ==nil && right ==nil {
		return true
		
	}else if left.Val !=right.Val{
		return false
	}
	in := compare(left.Right,right.Left)
	out := compare(left.Left,right.Right)
	return in&&out

}
// @lc code=end

