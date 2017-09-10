package tooffer

// Search 为面试题 3
func Search(array [][]int, number int) bool {
	rs, cs := len(array), len(array[0])
	r, c := 0, cs-1
	for r < rs && c >= 0 {
		if array[r][c] == number {
			return true
		} else if array[r][c] > number {
			c--
		} else {
			r++
		}
	}

	return false
}
