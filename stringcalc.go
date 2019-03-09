package stringcalc

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const defaultDelimiter = ","

var delimiterRegExp = regexp.MustCompile(`\[(.+?)]`)

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

func (sc StringCalc) parseInput(input string) []string {
	if input == "" {
		return nil
	}

	delimiters := sc.parseDelimiters(input)
	input = sc.cutOutDelimitersIfPresent(input)
	input = strings.ReplaceAll(input, "\n", defaultDelimiter)

	return sc.applyDelimiters(input, delimiters)
}

func (sc StringCalc) parseDelimiters(input string) []string {
	switch {
	case strings.HasPrefix(input, "//["):
		return sc.parseComplexDelimiters(input)
	case strings.HasPrefix(input, "//"):
		return sc.parseSimpleDelimiters(input)
	default:
		return nil
	}
}

func (sc StringCalc) parseSimpleDelimiters(input string) []string {
	delimiter := strings.TrimPrefix(input, "//")
	idx := strings.Index(delimiter, "\n")
	delimiter = delimiter[:idx]
	return []string{delimiter}
}

func (sc StringCalc) parseComplexDelimiters(input string) []string {
	var delimiters []string
	for _, match := range delimiterRegExp.FindAllStringSubmatch(input, -1) {
		delimiters = append(delimiters, match[1])
	}
	return delimiters
}

func (sc StringCalc) cutOutDelimitersIfPresent(input string) string {
	if !strings.HasPrefix(input, "//") {
		return input
	}

	idx := strings.Index(input, "\n")
	input = input[idx+1:]

	return input
}

func (sc StringCalc) applyDelimiters(line string, delimiters []string) []string {
	for _, delimiter := range delimiters {
		line = strings.ReplaceAll(line, delimiter, defaultDelimiter)
	}
	return strings.Split(line, defaultDelimiter)
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
