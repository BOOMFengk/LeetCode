/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */
//sum是记录从statidx开始 如果有负的 则不行
//tot是记录总的 如果是负的则说明跑不了一圈
// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	startidx, sum, tot := 0, 0, 0
	for i := range gas {
		temp := gas[i] - cost[i]
		sum += temp
		tot += temp
		if sum < 0 {
			startidx = i + 1
			sum = 0
		}

	}
	if tot < 0 {
		return -1
	}
	return startidx

}

// @lc code=end

