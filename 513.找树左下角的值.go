/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	depthMax,depthLeft:=0,0
	var trace func(*TreeNode,int)
	trace = func (root *TreeNode,depth int) {
		if root == nil {
			return
		}
		//同一层会先遍历左边如果左边有值就不会更新右侧结果
		if root.Left == nil && root.Right ==nil &&depth>depthMax{
			depthMax = depth
			depthLeft = root.Val
		}
		trace(root.Left,depth+1)
		trace(root.Right,depth+1)
		
	}

		trace(root,1)
		return depthLeft

}
// @lc code=end

