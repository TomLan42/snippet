/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			// record content
			temp := make([]byte, 0)
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp = append(temp, v)
			}

			// pop '['
			stack = stack[:len(stack)-1]

			// pop num
			// record num str len
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
			}
			//note: +1 is to compensate idx++
			num := stack[len(stack)-idx+1:]
			stack = stack[:len(stack)-idx+1]
			count, _ := strconv.Atoi(string(num))

			for i := 0; i < count; i++ {
				for j := len(temp) - 1; j >= 0; j-- {
					stack = append(stack, temp[j])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// @lc code=end

