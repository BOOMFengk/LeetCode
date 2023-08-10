/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
//第一轮循环把每个字符出现的位置用mp记录下来，走一遍之后就知道最后出现的位置在哪了
//然后第二轮循环就去找mp中记录的数和现在出现的角标位置重合的时候就到了这个字符串的结尾
func partitionLabels(s string)(ans []int) {
	mp := map[byte]int{}
	for idx,v := range s {
		ch:=byte(v)
		mp[ch]=idx	
	}
	startIdx,endIdx:=0,0
	str:=[]byte(s)
	curIdx:=mp[str[0]]
	for idx, ch := range str {
		curIdx = max(curIdx,mp[ch])
		if curIdx==idx{
			endIdx=idx
			ans=append(ans,endIdx-startIdx+1)
			startIdx=curIdx+1
		}
		
	}
	return
}
func max(i,j int) int {
	if i>j{
		return i
	}
	return j
	
}
// @lc code=end

