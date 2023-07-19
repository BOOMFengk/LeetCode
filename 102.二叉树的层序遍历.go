/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) (res[][]int) {
	if root ==nil{
		return
	}

	queue,level :=[]*TreeNode{root},0
	for len(queue) >0{
		//当前层有多少节点
		length := len(queue)
		res = append(res,[]int{})
		for i := 0; i < length; i++ {
			res[level] = append(res[level],queue[i].Val)
			
			if queue[i].Left !=nil{
				queue = append(queue,queue[i].Left)
			}
			if queue[i].Right !=nil{
				queue = append(queue,queue[i].Right)
			}
		}

		queue = queue[length:]
		level++
	}

	return res

}
// @lc code=end

