/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte) {

	reverse(s)
}

// recursive version
func reverse(s []byte) {
	if len(s) == 0 {
		return
	}
	reverseString(s[1:])

	temp := s[0]
	for i := 0; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	s[len(s)-1] = temp
	return
}

// swap version
func reverseSwap(s []byte) {

	left := 0
	right := len(s) - 1
	for left < right {
		s[left] = s[left] ^ s[right]
		s[right] = s[left] ^ s[right]
		s[left] = s[left] ^ s[right]
		left++
		right--
	}
	return
}

// @lc code=end

