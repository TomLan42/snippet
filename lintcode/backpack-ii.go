package lintcode

// backPack returns the maximum value can the items in A
// fill in the backpack m
func backPack2(m int, A []int, V []int) int {

	f := make([][]int, len(A)+1)
	// f[i][j] denotes the maximum value the first i items can fit
	// into the backpack of cap j
	// init
	for i := 0; i < len(A)+1; i++ {
		f[i] = make([]int, m+1)
	}

	for i := 1; i < len(A)+1; i++ {
		for j := 1; j < m+1; j++ {
			f[i][j] = f[i-1][j]
			if j-A[i-1] >= 0 {
				f[i][j] = max(f[i-1][j], f[i-1][j-A[i-1]]+V[i-1])
			}
		}
	}
	return f[len(A)][m]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
