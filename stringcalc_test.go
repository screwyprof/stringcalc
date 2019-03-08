package stringcalc_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/screwyprof/stringcalc"
)

func TestStringCalc_Add(t *testing.T) {
	t.Run("GivenEmptyInputZeroSumReturned", func(t *testing.T) {
		// arrange
		want := 0

		calc := stringcalc.StringCalc{}

		// act
		got, err := calc.Add("")

		// assert
		Ok(t, err)
		Equals(t, want, got)
	})

	t.Run("GivenOneNumberTheSameNumberReturned", func(t *testing.T) {
		// arrange
		want := 5

		calc := stringcalc.StringCalc{}

		// act
		got, err := calc.Add("5")

		// assert
		Ok(t, err)
		Equals(t, want, got)
	})

	t.Run("GivenTwoNumbersTheSumReturned", func(t *testing.T) {
		// arrange
		want := 7

		calc := stringcalc.StringCalc{}

		// act
		got, err := calc.Add("5,2")

		// assert
		Ok(t, err)
		Equals(t, want, got)
	})

	t.Run("GivenArbitraryNumbersTheSumReturned", func(t *testing.T) {
		// arrange
		want := 7

		calc := stringcalc.StringCalc{}

		// act
		got, err := calc.Add("3,2,1,0,1")

		// assert
		Ok(t, err)
		Equals(t, want, got)
	})

	t.Run("GivenNewLinesBetweenNumbersTheSumReturned", func(t *testing.T) {
		// arrange
		want := 6

		calc := stringcalc.StringCalc{}

		// act
		got, err := calc.Add("1\n2,3")

		// assert
		Ok(t, err)
		Equals(t, want, got)
	})

	t.Run("GivenInvalidInputAnErrorReturned", func(t *testing.T) {
		// arrange
		want := fmt.Errorf("invalid input given: lalaef,eff")

		calc := stringcalc.StringCalc{}

		// act
		_, err := calc.Add("lalaef,eff")

		// assert
		Equals(t, want, err)
	})

	t.Run("GivenACustomDelimiterTheSumReturned", func(t *testing.T) {
		// arrange
		want := 3

		calc := stringcalc.StringCalc{}

		// act
		got, err := calc.Add("//;\n1;2")

		// assert
		Ok(t, err)
		Equals(t, want, got)
	})

	t.Run("GivenNegativeNumbersAnErrorReturned", func(t *testing.T) {
		// arrange
		want := fmt.Errorf("negative numbers not allowed: -2,-4")

		calc := stringcalc.StringCalc{}

		// act
		_, err := calc.Add("1,-2,-4")

		// assert
		Equals(t, want, err)
	})
}

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	tb.Helper()
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
