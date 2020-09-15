package lco

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)

	if n == 0 {
		return false
	}

	m := len(matrix[0])

	r, c := n-1, 0
	for r >= 0 && c < m {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			r--
		} else if matrix[r][c] < target {
			c++
		}
	}

	return false
}
