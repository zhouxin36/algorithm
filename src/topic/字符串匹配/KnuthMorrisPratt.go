package 字符串匹配

func KMP(context, patStr string) int {
	res := computePrefix(patStr)
	var i, j int
	for i = 0; i < len(context) && j < len(patStr); {
		if j != -1 && context[i] != patStr[j] {
			j = res[j]
		} else {
			i++
			j++
		}
	}
	if j == len(patStr) {
		return i - j
	}
	return -1
}
func computePrefix(patStr string) []int {
	next := make([]int, len(patStr)+1)
	next[0] = -1
	j := -1
	for i := 0; i < len(patStr); i++ {
		if j != -1 && patStr[j] != patStr[i] {
			j = next[j]
		} else {
			i++
			j++
			next[i] = j
		}
	}

	return next
}
