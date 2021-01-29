package lintcode

// backPack returns the maximum space can the items in A
// fill in the backpack m
func backPack(m int, A []int) int {

	// f[i][j] denotes whether the first i items can exactly
	// fill in the backpack of capacity j
	f := make([][]bool, len(A)+1)
	// init
	for i := 0; i < len(A)+1; i++ {
		f[i] = make([]bool, m+1)
	}

	for i := 1; i < len(A)+1; i++ {
		for j := 1; j < m+1; j++ {
			f[i][j] = f[i-1][j]
			if j-A[i] >= 0 && f[i-1][j-A[i]] {
				f[i][j] = true
			}
		}
	}

	for j := m; j >= 1; j-- {
		if f[len(A)][j] {
			return j
		}
	}

	return 0

}
