package lco

func validateStackSequences(pushed []int, popped []int) bool {
	stack, i := make([]int, 0), 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) != 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}

func validateStackSequences2(pushed []int, popped []int) bool {
	i, j := 0, 0
	for _, v := range pushed {
		pushed[i] = v
		for i >= 0 && pushed[i] == popped[j] {
			j++
			i--
		}
		i++
	}
	return i == 0
}
