/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums,func(i,j int)bool{
		return abs(nums[i])<abs(nums[j])
	})
	
	for i:=len(nums)-1;i>0;i--{
		if k>0 &&nums[i]<0 {
			nums[i]*=-1
			k--
		}
	}
	if k%2==1{
		nums[0]*=-1
	}
	ans:=0
	for _, v := range nums {
		ans+=v
	}
	return ans

}

func abs(i int)int{
	if i<0{
		return -i
	}
	return i

}
// @lc code=end

