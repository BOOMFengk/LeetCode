/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) (result int) {
	for i := 1; i < len(prices); i++ {
		if prices[i] - prices[i-1]>0{
		result += prices[i] - prices[i-1]
	}
	
}
		return 
}
// @lc code=end

