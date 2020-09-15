package lc

func shortestSeq(big []int, small []int) []int {
	bigLen := len(big)
	smallLen := len(small)
	countMap := make(map[int]int)
	smallMap := make(map[int]int)
	bigMap := make(map[int]int)
	queue := make([]int, 0)
	result := []int{-1, -1}
	length := bigLen + 1

	for _, v := range small {
		smallMap[v] = 0
	}

	for i := 0; i < bigLen; i++ {
		key := big[i]
		if _, ok := smallMap[key]; ok {
			queue = append(queue, i)
			bigMap[key] = 0
			if _, ok := countMap[key]; ok {
				countMap[key] += 1
			} else {
				countMap[key] = 1
			}

			if len(bigMap) == smallLen {
				temp := 0
				for {
					temp = queue[0]
					queue = queue[1:]
					count := countMap[big[temp]] - 1
					countMap[big[temp]] = count
					if count == 0 {
						delete(bigMap, big[temp])
						break
					}
				}
				if length > i-temp+1 {
					length = i - temp + 1
					result[0] = temp
					result[1] = i
				}
			}
		}
	}

	if result[0] < 0 {
		return []int{}
	}
	return result
}
