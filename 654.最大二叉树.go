/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
 //大概思想也是每次去找那个最大值 然后去分更小的点去构建左子树和右子树
 //"这个过程可以视为在每次迭代中找到数组中的最大值，并用这个最大值作为根节点。然后，根据这个最大值将数组分割成两部分，左侧的元素构成左子树，右侧的元素构成右子树。这个过程不断递归进行，直到数组被分解完毕，此时二叉树也就构建完成。"
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}

	idx := searchMax(nums)
	root :=&TreeNode{nums[idx],nil,nil}
	
	leftTree,rightTree:=nums[:idx],nums[idx+1:]
	root.Left = constructMaximumBinaryTree(leftTree)
	root.Right = constructMaximumBinaryTree(rightTree)
	return root
}

func searchMax(nums []int)int{
	cnt := nums[0]
	maxIndex :=0
	for i,v :=range nums{
		if v>cnt{
		cnt=max(cnt,v)
		maxIndex = i
		}
	}
	return maxIndex
}

func max(i,j int)int{
	if i>j{
		return i
	}
	return j 
}
// @lc code=end

