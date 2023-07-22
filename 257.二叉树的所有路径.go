/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var trace func (*TreeNode,string)  
	trace = func (root *TreeNode,path string)  {
		path +=strconv.Itoa(root.Val)
		if root.Left==nil&&root.Right==nil{
			res = append(res,path)
		}
		if root.Left!=nil{
			trace(root.Left,path+"->")
		}
		if root.Right!=nil{
			trace(root.Right,path+"->")
		}
		
	}
	trace(root,"")
	return res

}
// @lc code=end

