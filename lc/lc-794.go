package lc

func validTicTacToe(board []string) bool {
	Xcount, Ocount := 0, 0
	for _, s := range board {
		for _, v := range s {
			if v == 'X' {
				Xcount++
			} else if v == 'O' {
				Ocount++
			}
		}
	}
	win := func(c uint8) bool {
		for i := 0; i < 3; i++ {
			if board[0][i] == c && board[1][i] == c && board[2][i] == c {
				return true
			}
			if board[i][0] == c && board[i][1] == c && board[i][2] == c {
				return true
			}
		}
		if board[0][0] == c && board[1][1] == c && board[2][2] == c {
			return true
		}
		if board[0][2] == c && board[1][1] == c && board[2][0] == c {
			return true
		}
		return false
	}
	if Xcount != Ocount && Ocount+1 != Xcount ||
		win('X') && Ocount+1 != Xcount ||
		win('O') && Ocount != Xcount {
		return false
	}
	return true
}
