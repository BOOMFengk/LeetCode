/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	ans := make([][]string, 0)
	var trace func(path []string,idx int)
	trace = func (path []string,idx int)  {
		if len(path)>4{
			return
		}
		if idx == len(s) &&len(path) == 4{
			temp:=make([]string, len(path))
			copy(temp,path)
			ans = append(ans,temp)
			return
		}
		for i := idx; i < len(s); i++ {
			temp := s[idx:i+1]
			if cnt,ok:=check(temp);ok{
				next:=append(path,cnt)
				trace(next,i+1)
			}
		}
	}
	trace([]string{},0)
	var result []string
	for i := range ans {
		temp:=strings.Join(ans[i],".")
		result = append(result,temp)
		
	}
	return result

}


func check(s string)(string,bool){
	if s[0] == '0' && len(s) != 1{
		return "",false
	}
	cnt,_:=strconv.Atoi(s)
	if cnt>=0&&cnt <= 255{
		return strconv.Itoa(cnt),true
	}
	return "",false
}
// @lc code=end

