package lc

func findLadders(beginWord string, endWord string, wordList []string) []string {
	var (
		canTranslate func(from string, to string) bool
		hasRoute     func(curWord string) bool
		visited      []bool
		result       []string
	)

	canTranslate = func(from string, to string) bool {
		if len(from) != len(to) {
			return false
		}
		count := 0
		for i := 0; i < len(from); i++ {
			if from[i] != to[i] {
				count++
			}
		}
		return count == 1
	}

	hasRoute = func(curWord string) bool {
		if curWord == endWord {
			return true
		}
		for i := 0; i < len(wordList); i++ {
			if visited[i] || !canTranslate(curWord, wordList[i]) {
				continue
			}
			visited[i] = true
			result = append(result, wordList[i])
			if hasRoute(wordList[i]) {
				return true
			}
			result = result[:len(result)-1]
			//visited[i]=false
		}
		return false
	}

	visited = make([]bool, len(wordList))
	result = make([]string, 0)

	for i := 0; i < len(wordList); i++ {
		visited[i] = false
	}
	result = append(result, beginWord)
	if hasRoute(beginWord) {
		return result
	}
	return []string{}
}
