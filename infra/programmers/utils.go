package programmers

import "strconv"

func toString(number uint64) string {
	return strconv.FormatInt(int64(number), 10)
}
