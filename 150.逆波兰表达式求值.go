/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack:= make([]int,0, len(tokens))
	for i := 0; i < len(tokens); i++ {
		val := tokens[i]
		switch val {
		case "+","-","*","/":
			a2:=stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a1:=stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result:=0

			switch val{
			case "+":
				result = a1+a2
			case "-":
				result = a1-a2
			case "*":
				result = a1*a2
			case "/":
				result = a1/a2
			}
			stack = append(stack,result)
		default:
			v,_ :=strconv.Atoi(val)
			stack = append(stack,v)

		}
	}
	return stack[0]

}
// @lc code=end

