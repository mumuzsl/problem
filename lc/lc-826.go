package lc

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	difficultyLen := len(difficulty)
	workerLen := len(worker)
	ans := 0
	var curMaxPro int
	for j := 0; j < workerLen; j++ {
		curMaxPro = 0
		for i := 0; i < difficultyLen; i++ {
			if difficulty[i] <= worker[j] && profit[i] > curMaxPro {
				curMaxPro = profit[i]
			}
		}
		ans += curMaxPro
	}
	return ans
}
