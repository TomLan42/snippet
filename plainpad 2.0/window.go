package main

import "math"

func minWindow(s string, t string) string {
	win := make(map[byte]int)
	need := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	right := 0

	match := 0
	start := 0
	end := 0
	min := math.MaxInt64

	var c byte
	for right < len(s) {

		c = s[right]
		right++

		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				// indicate match of a single character
				match++
			}
		}

		// all match
		for match == len(need) {
			if right-left < min {
				min = right - min
				start = left
				end = right
			}

			c = s[left]
			left++
			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}

	if min == math.MaxInt64 {
		return ""
	}

	return s[start:end]
}
