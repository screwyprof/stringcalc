package stringcalc

import (
	"fmt"
	"strconv"
	"strings"
)

type StringCalc struct {
}

func (sc StringCalc) Add(nums string) (int, error) {
	if nums == "" {
		return 0, nil
	}

	numbers := sc.parseInput(nums)
	if len(numbers) < 1 {
		return 0, nil
	}

	if len(numbers) == 1 {
		num, err := strconv.Atoi(numbers[0])
		if err != nil {
			return 0, fmt.Errorf("invalid input given: %s", nums)
		}
		return num, nil
	}

	var sum int
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			return 0, fmt.Errorf("invalid input given: %s", nums)
		}
		sum += num
	}

	return sum, nil
}

func (sc StringCalc) parseInput(nums string) []string {
	lines, delimiter := sc.ripOffDelimiter(strings.Split(nums, "\n"))

	var numbers []string
	for _, line := range lines {
		figs := strings.Split(line, delimiter)
		numbers = append(numbers, figs...)
	}

	return numbers
}

func (sc StringCalc) ripOffDelimiter(lines []string) ([]string, string) {
	const defaultDelimiter = ","
	if !strings.HasPrefix(lines[0], "//") {
		return lines, defaultDelimiter
	}

	delimiter := strings.TrimPrefix(lines[0], "//")
	lines = append(lines[:0], lines[1:]...)

	return lines, delimiter
}
