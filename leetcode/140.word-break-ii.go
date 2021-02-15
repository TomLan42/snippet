/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	// backtracking
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	result := make([][]string, 0)
	// backtrack all split word combo
	backtrack(s, 0, []string{}, &result, dict)

	finalResult := make([]string, 0)
	for _, words := range result {
		str := ""
		for _, word := range words {
			str += word + " "
		}
		finalResult = append(finalResult, str[:len(str)-1])
	}

	return finalResult

}

func backtrack(s string, idx int, cur []string, result *[][]string, dict map[string]bool) {
	start, end := idx, len(s)

	if start == end {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*result = append(*result, temp)
	}

	for i := start; i < end; i++ {
		if inDict(dict, s[start:i+1]) {
			backtrack(s, i+1, append(cur, s[start:i+1]), result, dict)
		}
	}

}

func inDict(dict map[string]bool, target string) bool {
	if _, found := dict[target]; found {
		return true
	}
	return false
}

// @lc code=end

