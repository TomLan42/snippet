package main

func editSpace(s string, t string) int {

	// state init
	matrix := make([][]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		matrix[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i] == t[j] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = matrix[i-1][j-1] + 1
			}

			if matrix[i][j] > matrix[i-1][j]+1 {
				matrix[i][j] = matrix[i-1][j] + 1
			}

			if matrix[i][j] > matrix[i][j-1]+1 {
				matrix[i][j] = matrix[i][j-1] + 1
			}

		}
	}

	return matrix[len(s)][len(t)]
}
