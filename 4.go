package tooffer

// EscapeSpace 对空格进行转义
func EscapeSpace(bs []byte) []byte {
	i := len(bs) - 1
	spaceCount := 0
	for _, c := range bs {
		if c == ' ' {
			spaceCount++
		}
	}
	spaces := make([]byte, spaceCount*2)
	bs = append(bs, spaces...)
	j := len(bs) - 1
	for j > i {
		if bs[i] == ' ' {
			bs[j-2] = '%'
			bs[j-1] = '2'
			bs[j] = '0'
			j -= 3
			i--
		} else {
			bs[j] = bs[i]
			j--
			i--
		}
	}

	return bs
}
