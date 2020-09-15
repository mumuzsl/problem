package lc

import "fmt"

func solveSudoku(board [][]byte) {
	const (
		N   = 9
		NEG = (1 << 9) - 1
	)
	var (
		rows  = make([]int, N)
		cols  = make([]int, N)
		cells = make([]int, N)
	)
	var (
		dfs               func(cnt int) bool
		getNext           func() (x, y int)
		getPossibleStatus func(x, y int) int
		fillNum           func(x, y, n int, fillFlag bool)
		count             func(x int) int
	)

	//golang中不能像C++代码那样直接用"~"实现各位取反。
	//但可以通过与相同位数且每位都是1的二进制数进行"^"(异或)运算从而实现各位取反。
	//这里的二进制数共有9位，所以常数NEG=(1 << 9) - 1。
	//因为1<<9的二进制数为1000000000，减1之后可以得到二进制为8位且每位为1的整数。
	getPossibleStatus = func(x, y int) int {
		return (rows[x] | cols[y] | cells[(x/3)*3+(y/3)]) ^ NEG
	}

	//计算整数x中1的个数
	count = func(x int) int {
		c := 0
		for x != 0 {
			x = x & (x - 1)
			c++
		}
		return c
	}

	getNext = func() (x, y int) {
		var cur int
		minCnt := 10

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if board[i][j] != '.' {
					continue
				}
				cur = count(getPossibleStatus(i, j))
				if cur >= minCnt {
					continue
				}
				x, y = i, j
				minCnt = cur
			}
		}
		return x, y
	}

	fillNum = func(x, y, n int, fillFlag bool) {
		pick := 1 << n
		if fillFlag {
			rows[x] |= pick
			cols[y] |= pick
			cells[(x/3)*3+(y/3)] |= pick
		} else {
			rows[x] ^= pick
			cols[y] ^= pick
			cells[(x/3)*3+(y/3)] ^= pick
		}
	}

	dfs = func(cnt int) bool {
		if cnt == 0 {
			return true
		}
		x, y := getNext()
		bits := getPossibleStatus(x, y)
		for i := 0; i < N; i++ {
			if (bits & (1 << i)) == 0 {
				continue
			}
			//fmt.Printf("%9b, %9b, %9b, %9b ", bits, rows[x], cols[y], cells[(x/3)*3+(y/3)])
			//fmt.Println(board[x][y], byte(i+'1'), i, x, y, cnt)
			fillNum(x, y, i, true)
			board[x][y] = byte(i + '1')
			if dfs(cnt - 1) {
				return true
			}
			board[x][y] = '.'
			fillNum(x, y, i, false)
		}
		return false
	}

	cnt := 0
	var n int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == '.' {
				cnt++
				continue
			}
			n = 1 << int(board[i][j]-'1')
			rows[i] |= n
			cols[j] |= n
			cells[(i/3)*3+(j/3)] |= n
		}
	}
	//fmt.Println(board, cnt)
	//fmt.Println(rows, cols, cells)
	fmt.Println(dfs(cnt))
}
