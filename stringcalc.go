package stringcalc

import (
	"strconv"
	"strings"
)

type StringCalc struct {
}

func (sc StringCalc) Add(nums string) int {
	numbers := sc.parseNumbers(nums)

	if len(numbers) < 1 {
		return 0
	}

	if len(numbers) == 1 {
		num, _ := strconv.Atoi(numbers[0])
		return num
	}

	var sum int
	for _, number := range numbers {
		num, _ := strconv.Atoi(number)
		sum += num
	}

	return sum
}

func (sc StringCalc) parseNumbers(nums string) []string {
	var numbers []string

	lines := strings.Split(nums, "\n")
	for _, line := range lines {
		figs := strings.Split(line, "")
		numbers = append(numbers, figs...)
	}

	return numbers
}
