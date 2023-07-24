/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

 //有点分而划之的感觉  每次都会找出中序中的左右根，然后把中序的左边去传递继续进行分割左右
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder)==0{
		return nil
	}

	root := &TreeNode{postorder[len(postorder)-1],nil,nil}
	
	rootIdx:=0
	for i:=0; i<len(inorder);i++{
		if inorder[i]==root.Val{
		rootIdx = i
		break
	}
	}

	leftInorder := inorder[:rootIdx]
	rightInorder := inorder[rootIdx+1:]

	//这里就是去找左子树的根 
	leftPostOrder := postorder[:len(leftInorder)]
	rightPostOrder := postorder[len(leftInorder):len(postorder)-1]

	//将中序的左子树，后序的长度传进去，大概就构成了令一个小的部分子树，从而一直分下去
	root.Left = buildTree(leftInorder,leftPostOrder)
	root.Right = buildTree(rightInorder,rightPostOrder)
	return root
}
// @lc code=end

