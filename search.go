package tooffer

import (
	"bytes"
	// "fmt"
)

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

// SpaceEscape 对空格进行转义
func SpaceEscape(s string) string {
	var buf bytes.Buffer

	for _, c := range s {
		if c == ' ' {
			buf.WriteString("%20")
		} else {
			buf.WriteByte(byte(c))
		}
	}

	return buf.String()
}
