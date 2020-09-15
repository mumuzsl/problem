package lco

func exist(board [][]byte, word string) bool {
	rLen := len(board)
	if rLen == 0 {
		return false
	}
	cLen := len(board[0])
	wordLen := len(word)

	flag := make([]bool, rLen*cLen)
	var search func(r, c, n int) bool
	search = func(r, c, n int) bool {
		if (r >= rLen || r < 0) || (c >= cLen || c < 0) || flag[(r*cLen+c)] || board[r][c] != word[n] {
			return false
		}
		if n == wordLen-1 {
			return true
		}
		flag[(r*cLen + c)] = true
		res := search(r+1, c, n+1) || search(r, c+1, n+1) || search(r-1, c, n+1) || search(r, c-1, n+1)
		flag[(r*cLen + c)] = false
		return res
	}

	for i := 0; i < rLen*cLen; i++ {
		if search(i/cLen, i%cLen, 0) {
			return true
		}
	}

	return false
}
