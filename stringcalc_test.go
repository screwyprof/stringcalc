package stringcalc_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
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
		{"GivenEmptyInputZeroSumReturned", "", checks(validCase(0))},
		{"GivenOneNumberTheSameNumberReturned", "5", checks(validCase(5))},
		{"GivenTwoNumbersTheSumReturned", "5,2", checks(validCase(7))},
		{"GivenArbitraryNumbersTheSumReturned", "3,2,1,0,1", checks(validCase(7))},
		{"GivenNewLinesBetweenNumbersTheSumReturned", "1\n2,3", checks(validCase(6))},
		{"GivenInvalidInputAnErrorReturned", "lalaef,eff", checks(hasError(errInvalidInput))},
		{"GivenACustomDelimiterTheSumReturned", "//;\n1;2", checks(validCase(3))},
		{"GivenNegativeNumbersAnErrorReturned", "1,-2,-4", checks(hasError(errNegativesNotAllowed))},
		{"GivenANumberGreaterThan1000ItIsNotSummed", "2,1001", checks(validCase(2))},
		{"GivenALengthyDelimiterTheSumReturned", "//[***]\n1***2***3", checks(validCase(6))},
		{"GivenAFewDelimitersTheSumReturned", "//[*][%]\n1*2%3", checks(validCase(6))},
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

func hasError(want error) check {
	return func(t *testing.T, got int, err error) {
		t.Helper()
		equals(t, want, err)
	}
}

func validCase(want int) check {
	return func(t *testing.T, got int, err error) {
		t.Helper()
		ok(t, err)
		equals(t, want, got)
	}
}

// Assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	tb.Helper()
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
