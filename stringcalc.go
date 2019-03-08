package stringcalc

import (
	"errors"
	"fmt"
	"regexp"
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
		return 0, errors.New("invalid input given: " + nums)
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

	lines := strings.Split(nums, "\n")

	delimiters := []string{","}
	if strings.HasPrefix(lines[0], "//") {
		delimiters = sc.parseDelimiters(lines[0])
		lines = append(lines[:0], lines[1:]...)
	}

	var numbers []string
	for _, line := range lines {
		numbers = append(numbers, sc.applyDelimiters(line, delimiters)...)
	}

	return numbers
}

func (sc StringCalc) applyDelimiters(line string, delimiters []string) []string {
	for _, delimiter := range delimiters {
		line = strings.ReplaceAll(line, delimiter, ",")
	}
	return strings.Split(line, ",")
}

func (sc StringCalc) parseDelimiters(delimiterStr string) []string {
	// one simple delimiter
	if !strings.HasPrefix(delimiterStr, "//[") {
		delimiterStr = strings.TrimPrefix(delimiterStr, "//")
		return []string{delimiterStr}
	}

	var re = regexp.MustCompile(`\[(.+?)]`)

	var delimiters []string
	for _, match := range re.FindAllStringSubmatch(delimiterStr, -1) {
		delimiters = append(delimiters, match[1])
	}

	return delimiters
}

func (sc StringCalc) strSliceToIntSlice(strNums []string) ([]int, error) {
	var numbers []int
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return nil, errors.New("invalid number given: " + strNum)
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
		return errors.New("negative numbers not allowed: " + sc.numbersToString(negativeNumbers))
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
