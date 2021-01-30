/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {

	// equal to find the lower bound
	left := 1
	right := n

	for left+1 < right {
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	if isBadVersion(left) {
		return left
	}
	if isBadVersion(right) {
		return right
	}
	return left

}

// @lc code=end

