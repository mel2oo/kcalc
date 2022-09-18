package conv

import "strconv"

func Int2Bin(i int64, length int) string {
	return align(strconv.FormatInt(i, 2), length)
}

func align(str string, length int) (s string) {
	l1 := length - len(str)
	for l1 > 0 {
		s += "0"
		l1--
	}

	return s + str
}
