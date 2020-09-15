package lc

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	ans := 0
	var cur int
	var dfs func(i int)
	dfs = func(i int) {
		if manager[i] == -1 {
			return
		}
		cur += informTime[manager[i]]
		dfs(manager[i])
	}
	for i, v := range informTime {
		if v == 0 {
			cur = 0
			dfs(i)
			if cur > ans {
				ans = cur
			}
		}
	}
	return ans
}

func numOfMinutes2(n int, headID int, manager []int, informTime []int) int {
	ans := 0
	var cur, j int
	for i, v := range informTime {
		if v == 0 {
			cur = 0
			j = i
			for manager[j] != -1 {
				cur += informTime[manager[j]]
				j = manager[j]
			}
			if cur > ans {
				ans = cur
			}
		}
	}
	return ans
}

func numOfMinutes3(n int, headID int, manager []int, informTime []int) int {
	table := make([][]int, n)
	ans := 0
	var j int
	for i := 0; i < n; i++ {
		j = manager[i]
		if j == -1 {
			continue
		}
		if table[j] == nil {
			table[j] = make([]int, 0)
		}
		table[j] = append(table[j], i)
	}
	var dfs func(index, sum int)
	dfs = func(index, sum int) {
		if table[index] == nil {
			if sum > ans {
				ans = sum
			}
			return
		}
		for _, v := range table[index] {
			dfs(v, sum+informTime[index])
		}
	}
	dfs(headID, 0)
	return ans
}
