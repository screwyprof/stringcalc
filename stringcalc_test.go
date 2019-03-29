package stringcalc_test

import (
	"fmt"
	"testing"

	"github.com/screwyprof/stringcalc"
)

var (
	errInvalidInput        = fmt.Errorf("invalid input given: lalaef,eff")
	errNegativesNotAllowed = fmt.Errorf("negative numbers not allowed: -2,-4")
)

func TestStringCalc_Add(t *testing.T) {
	checks := func(cs ...check) []check { return cs }

	cases := []struct {
		name   string
		input  string
		checks []check
	}{
		{"GivenEmptyInputZeroSumReturned", "", checks(expect(0))},
		{"GivenOneNumberTheSameNumberReturned", "5", checks(expect(5))},
		{"GivenTwoNumbersTheSumReturned", "5,2", checks(expect(7))},
		{"GivenArbitraryNumbersTheSumReturned", "3,2,1,0,1", checks(expect(7))},
		{"GivenNewLinesBetweenNumbersTheSumReturned", "1\n2,3", checks(expect(6))},
		{"GivenInvalidInputAnErrorReturned", "lalaef,eff", checks(expectError(errInvalidInput))},
		{"GivenACustomDelimiterTheSumReturned", "//;\n1;2", checks(expect(3))},
		{"GivenNegativeNumbersAnErrorReturned", "1,-2,-4", checks(expectError(errNegativesNotAllowed))},
		{"GivenANumberGreaterThan1000ItIsNotSummed", "2,1001", checks(expect(2))},
		{"GivenALengthyDelimiterTheSumReturned", "//[***]\n1***2***3", checks(expect(6))},
		{"GivenAFewDelimitersTheSumReturned", "//[*][%]\n1*2%3", checks(expect(6))},
	}

	calc := stringcalc.StringCalc{}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calc.Add(tc.input)
			for _, ch := range tc.checks {
				ch(t, got, err)
			}
		})
	}
}

type check func(t *testing.T, got int, err error)

func expectError(want error) check {
	return func(t *testing.T, got int, err error) {
		t.Helper()
		equals(t, want, err)
	}
}

func expect(want int) check {
	return func(t *testing.T, got int, err error) {
		t.Helper()
		ok(t, err)
		equals(t, want, got)
	}
}
