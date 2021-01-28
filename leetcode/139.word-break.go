/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {

	if len(s) == 0 {
		return true
	}

	max, dict := maxLen(wordDict)

	// f[i] means the first i chars conforms to whatever
	// the rool is
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 1; i <= len(s); i++ {
		l := 0
		if i-max > 0 {
			l = i - max
		}
		for j := l; j < i; j++ {
			if inDict(s[j:i], dict) && f[j] {
				f[i] = true
				break
			}
		}
	}

	return f[len(s)]

}

func maxLen(wordDict []string) (int, map[string]bool) {
	max := 0
	dict := make(map[string]bool)
	for _, str := range wordDict {
		if len(str) > max {
			max = len(str)
		}
		dict[str] = true
	}

	return max, dict
}

func inDict(str string, dict map[string]bool) bool {
	_, found := dict[str]
	return found
}

// @lc code=end

