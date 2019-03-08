package stringcalc

import (
	"fmt"
	"strconv"
	"strings"
)

type StringCalc struct{}

func (sc StringCalc) Add(nums string) (int, error) {
	strNumbers := sc.parseInput(nums)
	if len(strNumbers) < 1 {
		return 0, nil
	}

	numbers, err := sc.strSliceToIntSlice(strNumbers)
	if err != nil {
		return 0, fmt.Errorf("invalid input given: " + nums)
	}

	if negErr := sc.ensureNoNegativeNumbersGiven(numbers); negErr != nil {
		return 0, negErr
	}

	return sc.sumNumbers(numbers), nil
}

func (sc StringCalc) parseInput(nums string) []string {
	if nums == "" {
		return nil
	}

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

func (sc StringCalc) strSliceToIntSlice(strNums []string) ([]int, error) {
	var numbers []int
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return nil, fmt.Errorf("invalid number given: %s", strNum)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func (sc StringCalc) ensureNoNegativeNumbersGiven(nums []int) error {
	var negativeNumbers []int
	for _, num := range nums {
		if num < 0 {
			negativeNumbers = append(negativeNumbers, num)
		}
	}

	if len(negativeNumbers) > 0 {
		return fmt.Errorf("negative numbers not allowed: " + sc.numbersToString(negativeNumbers))
	}

	return nil
}

func (sc StringCalc) numbersToString(numbers []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", ",", -1), "[]")
}

func (sc StringCalc) sumNumbers(nums []int) int {
	var sum int
	for _, num := range nums {
		if num > 1000 {
			continue
		}
		sum += num
	}

	return sum
}
