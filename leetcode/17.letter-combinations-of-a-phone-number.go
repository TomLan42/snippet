/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	phone := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	if len(digits) == 0 {
		return []string{}
	}

	return backtrack(digits, phone)

}

func backtrack(digits string, phone map[string]string) []string {

	if len(digits) == 0 {
		return []string{""}
	}

	results := make([]string, 0)

	for i := 0; i < len(phone[digits[0:1]]); i++ {
		lists := backtrack(digits[1:], phone)
		for j := 0; j < len(lists); j++ {
			cur := phone[digits[0:1]][i : i+1]
			cur = cur + lists[j]
			results = append(results, cur)
		}
	}

	return results
}

// @lc code=end

