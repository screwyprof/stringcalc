package stringcalc

import "strconv"

type StringCalc struct {
}

func (sc StringCalc) Add(nums string) int {

	num, _ := strconv.Atoi(nums)

	return num
}
