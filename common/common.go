package common

import "strconv"

func ToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func Transpose(s [][]int) [][]int {
	xl := len(s[0])
	yl := len(s)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = s[j][i]
		}
	}
	return result
}

func Reverse(s []int) []int {
	slen := len(s)
	output := make([]int, slen)

	for i, n := range s {
		j := slen - i - 1

		output[j] = n
	}

	return output
}
