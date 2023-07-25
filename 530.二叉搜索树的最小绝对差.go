/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	ans := math.MaxInt32
	var trace func(root *TreeNode)
	trace = func (root *TreeNode)  {
		if root == nil{
			return
		}
		trace(root.Left)

		if prev!=nil{
			ans = min(ans,root.Val-prev.Val)
		}
		prev = root
		trace(root.Right)
	}
		trace(root)

		return ans
}

func min(i,j int)int{
	if i<j{
		return i
	}
	return  j
	
}
// @lc code=end

