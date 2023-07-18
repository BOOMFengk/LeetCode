/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	str := []byte(s)
	for i := 0; i < len(s); i = i + 2*k {
		if i+k < len(str) {
			reverseString(str[i : i+k])
		} else {
			reverseString(str[i:])
		}
	}
	return string(str)
}
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return
}

// @lc code=end

