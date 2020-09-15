package lco

func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)
	cols := lenP + 1
	flag := make([]bool, (lenS+1)*(lenP+1))

	for i := 0; i <= lenS; i++ {
		for j := 0; j <= lenP; j++ {
			if j == 0 {
				flag[i*cols+j] = i == 0
			} else {
				if p[j-1] != '*' {
					if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
						flag[i*cols+j] = flag[(i-1)*cols+(j-1)]
					}
				} else {
					if j >= 2 {
						flag[i*cols+j] = flag[i*cols+j] || flag[i*cols+(j-2)]
					}
					if i >= 1 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
						flag[i*cols+j] = flag[i*cols+j] || flag[(i-1)*cols+j]
					}
				}
			}
		}
	}

	return flag[lenS*cols+lenP]
}
