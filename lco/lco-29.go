package lco

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := make([]int, len(matrix)*len(matrix[0]))
	count := 0
	for {
		for i := l; i <= r; i++ {
			//res = append(res, matrix[t][i])
			res[count] = matrix[t][i]
			count++
		}
		if t++; t > b {
			break
		}
		for i := t; i <= b; i++ {
			//res = append(res, matrix[i][r])
			res[count] = matrix[i][r]
			count++
		}
		if r--; l > r {
			break
		}
		for i := r; i >= l; i-- {
			//res = append(res, matrix[b][i])
			res[count] = matrix[b][i]
			count++
		}
		if b--; t > b {
			break
		}
		for i := b; i >= t; i-- {
			//res = append(res, matrix[i][l])
			res[count] = matrix[i][l]
			count++
		}
		if l++; l > r {
			break
		}
	}
	return res
}
