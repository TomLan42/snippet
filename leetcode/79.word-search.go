/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {

	// bfs approach?
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				temp := board[i][j]
				board[i][j] = '0'
				if backtrack(i, j+1, word[1:], board) || backtrack(i, j-1, word[1:], board) || backtrack(i+1, j, word[1:], board) || backtrack(i-1, j, word[1:], board) {
					return true
				}
				board[i][j] = temp
			}
		}
	}
	return false
}

func backtrack(i, j int, word string, board [][]byte) bool {

	if len(word) == 0 {
		return true
	}

	// boundary check
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
		return false
	}

	if board[i][j] == word[0] {
		temp := board[i][j]
		board[i][j] = '0'
		if backtrack(i, j+1, word[1:], board) || backtrack(i, j-1, word[1:], board) || backtrack(i+1, j, word[1:], board) || backtrack(i-1, j, word[1:], board) {
			return true
		}
		board[i][j] = temp
	}
	return false
}

// @lc code=end

