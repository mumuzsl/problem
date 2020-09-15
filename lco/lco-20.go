package lco

import "fmt"

//有限状态自动机
func isNumber(s string) bool {
	states := []map[int8]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4}, // 0. start with 'blank'
		{'d': 2, '.': 4},                 // 1. 'sign' before 'e'
		{'d': 2, '.': 3, 'e': 5, ' ': 8}, // 2. 'digit' before 'dot'
		{'d': 3, 'e': 5, ' ': 8},         // 3. 'digit' after 'dot'
		{'d': 3},                         // 4. 'digit' after 'dot' (‘blank’ before 'dot')
		{'s': 6, 'd': 7},                 // 5. 'e'
		{'d': 7},                         // 6. 'sign' after 'e'
		{'d': 7, ' ': 8},                 // 7. 'digit' after 'e'
		{' ': 8},                         // 8. end with 'blank'
	}
	p := 0
	var t int8
	for _, c := range s {
		t = '?'
		if c >= '0' && c <= '9' {
			t = 'd'
		} else if c == '+' || c == '-' {
			t = 's'
		} else if c == 'e' || c == 'E' {
			t = 'e'
		} else if c == '.' || c == ' ' {
			t = int8(c)
		}
		if _, ok := states[p][t]; !ok {
			return false
		}
		p = states[p][t]
	}

	return p == 2 || p == 3 || p == 7 || p == 8
}

//剑指offer的解法
func isNumber_2(s string) bool {
	if len(s) == 0 {
		return false
	}

	// 整数的格式可以用[+|-]B表示, 其中B为无符号整数
	var scanInteger func(s string, index int) (bool, int)
	var scanUnsignedInteger func(s string, index int) (bool, int)
	scanUnsignedInteger = func(s string, index int) (bool, int) {
		befor := index
		for index != len(s) && s[index] >= '0' && s[index] <= '9' {
			index++
		}
		return index > befor, index
	}

	scanInteger = func(s string, index int) (bool, int) {
		if index != len(s) && (s[index] == '+' || s[index] == '-') {
			index++
		}
		return scanUnsignedInteger(s, index)
	}

	var index int
	var numeric bool
	for index < len(s) && s[index] == ' ' {
		index++
	}

	if index < len(s) {
		numeric, index = scanInteger(s, index)
	}
	fmt.Println(numeric, index)
	// 如果出现'.'，接下来是数字的小数部分
	if index < len(s) && s[index] == '.' {

		index++

		// 下面一行代码用||的原因：
		// 1. 小数可以没有整数部分，例如.123等于0.123；
		// 2. 小数点后面可以没有数字，例如233.等于233.0；
		// 3. 当然小数点前面和后面可以有数字，例如233.666
		//tm, index :=
		tm, ti := scanUnsignedInteger(s, index)
		numeric, index = tm || numeric, ti
	}

	// 如果出现'e'或者'E'，接下来跟着的是数字的指数部分
	if index < len(s) && (s[index] == 'e' || s[index] == 'E') {
		index++

		// 下面一行代码用&&的原因：
		// 1. 当e或E前面没有数字时，整个字符串不能表示数字，例如.e1、e1；
		// 2. 当e或E后面没有整数时，整个字符串不能表示数字，例如12e、12e+5.4
		tm, ti := scanInteger(s, index)
		fmt.Println(tm, ti)
		numeric, index = numeric && tm, ti
	}

	//字符串结尾有空格，可以返回true
	for index < len(s) && s[index] == ' ' {
		index++
	}

	fmt.Println(len(s), index)
	return numeric && index == len(s)
}

//不能通过
func isNumber2(s string) bool {
	isNum := func(v int8) bool { return v >= '0' && v <= '9' }

	table := make(map[int8]int, 5)
	table['+'] = 0
	table['-'] = 0
	table['.'] = 0
	table['E'] = 0
	table['e'] = 0

	var a func(v int8, i int) bool
	a = func(v int8, i int) bool {
		if i >= len(s) {
			return true
		}
		if (v == '0' && !isNum(int8(s[i]))) && v != int8(s[i]) {
			return false
		}
		switch v {
		case ' ':
			a(int8(s[i+1]), i+1)
		case '+', '-':
			table[v]++
			a('0', i+1)
			if table[v] == 1 {
				a('.', i+1)
			}
		case '.':
			table[v]++
			if table[v] == 1 {
				a('0', i+1)
			}
		case 'e', 'E':
			a('+', i+1)
			a('-', i+1)
			a('0', i+1)
		default:
			if isNum(v) {
				a('.', i+1)
				a('e', i+1)
				a('0', i+1)
			}
		}
		return false
	}

	return a(int8(s[0]), 0)
}
