/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) (ans []string ){
	if digits==""{
		return nil
	}
	var mp map[string]string = map[string]string{
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    }
	var trace func(idx int,cur string)
	trace = func (idx int,cur string)  {
		if len(cur)==len(digits){
			ans = append(ans,cur)
			return
		}

			chs := mp[string(digits[idx])]
			for _, v := range chs {
				trace(idx+1,cur+string(v))
				
			}
			
		
	}
	trace(0,"")
	return

}
// @lc code=end

