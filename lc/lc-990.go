package lc

func equationsPossible(equations []string) bool {
	parent := make([]int, 26)

	var union func(a, b int)
	var find func(x int) int
	union = func(a, b int) {
		parent[find(a)] = find(b)
	}
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}

	for i := 0; i < 26; i++ {
		parent[i] = i
	}
	var a, b int
	for _, v := range equations {
		if v[1] == '=' {
			a, b = int(v[0]-'a'), int(v[3]-'a')
			union(a, b)
		}
	}
	for _, v := range equations {
		if v[1] == '!' {
			a, b = int(v[0]-'a'), int(v[3]-'a')
			if find(a) == find(b) {
				return false
			}
		}
	}

	return true
}
