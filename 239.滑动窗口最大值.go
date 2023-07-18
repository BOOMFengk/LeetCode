/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	q :=make([]int, 0)
	var push func(x int)
	push = func (x int)  {
		for len(q)>0 &&q[len(q)-1]<x{
			q = q[:len(q)-1]
		}
		q = append(q,x)
	}

	for i := 0; i < k; i++ {
		push(nums[i])
	}
	ans[0]=q[0]

	for i := 1; i < len(nums)-k + 1; i++ {
		if nums[i-1]==q[0]{
			q=q[1:]
		}
		push(nums[i+k-1])
		
		ans[i]=q[0]
	}

	return ans

}
// @lc code=end

