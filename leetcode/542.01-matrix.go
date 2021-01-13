/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	q := make([][]int, 0)
	// transfer matrix
	// record 0 and mark non-zero
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}
	// start searching
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(q) != 0 {
		point := q[0]
		q = q[1:]

		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				q = append(q, []int{x, y})
			}
		}
	}
	return matrix
}

// @lc code=end

