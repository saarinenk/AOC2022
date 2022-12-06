package common

import "strconv"

func ToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
