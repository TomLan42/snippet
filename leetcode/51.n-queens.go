/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	col := make([]bool, n)
	row := make([]int, 0)
	dia1 := make([]bool, 2*n-1)
	dia2 := make([]bool, 2*n-1)
	res := [][]string{}

	putQueen(n, 0, &col, &dia1, &dia2, &row, &res)
	return res
}

func putQueen(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *[][]string) {
	if index == n {
		*res = append(*res, generateBoard(n, row))
	}

	for i := 0; i < n; i++ {
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i] = true
			(*dia2)[index-i+n-1] = true
			putQueen(n, index+1, col, dia1, dia2, row, res)
			(*col)[i] = false
			(*dia1)[index+i] = false
			(*dia2)[index-i+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}

	return
}

func generateBoard(n int, row *[]int) []string {
	board := []string{}
	res := ""
	for i := 0; i < n; i++ {
		res += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, res)
	}
	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)
	}
	return board
}

// @lc code=end

