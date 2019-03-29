package stringcalc_test

import (
	"testing"

	"github.com/screwyprof/stringcalc"
)

func TestStringCalc(t *testing.T) {
	spec(t, "GivenEmptyInputZeroSumReturned")(
		given(stringcalc.StringCalc{}),
		when(""),
		then(0),
	)

	spec(t, "GivenOneNumberTheSameNumberReturned")(
		given(stringcalc.StringCalc{}),
		when("5"),
		then(5),
	)

	spec(t, "GivenArbitraryNumbersTheSumReturned")(
		given(stringcalc.StringCalc{}),
		when("3,2,1,0,1"),
		then(7),
	)

	spec(t, "GivenNewLinesBetweenNumbersTheSumReturned")(
		given(stringcalc.StringCalc{}),
		when("1\n2,3"),
		then(6),
	)

	spec(t, "GivenInvalidInputAnErrorReturned")(
		given(stringcalc.StringCalc{}),
		when("lalaef,eff"),
		thenFailWith(errInvalidInput),
	)

	spec(t, "GivenACustomDelimiterTheSumReturned")(
		given(stringcalc.StringCalc{}),
		when("//;\n1;2"),
		then(3),
	)

	spec(t, "GivenNegativeNumbersAnErrorReturned")(
		given(stringcalc.StringCalc{}),
		when("1,-2,-4"),
		thenFailWith(errNegativesNotAllowed),
	)

	spec(t, "GivenANumberGreaterThan1000ItIsNotSummed")(
		given(stringcalc.StringCalc{}),
		when("2,1001"),
		then(2),
	)

	spec(t, "GivenALengthyDelimiterTheSumReturned")(
		given(stringcalc.StringCalc{}),
		when("//[***]\n1***2***3"),
		then(6),
	)

	spec(t, "GivenAFewDelimitersTheSumReturned")(
		given(stringcalc.StringCalc{}),
		when("//[*][%]\n1*2%3"),
		then(6),
	)
}

type givenFn func() stringcalc.StringCalc
type whenFn func(calc stringcalc.StringCalc) (int, error)
type thenFn func(t testing.TB) func(got int, err error)
type specFn func(given givenFn, when whenFn, then thenFn)

func spec(t *testing.T, name string) specFn {
	return func(given givenFn, when whenFn, then thenFn) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Helper()
			then(t)(when(given()))
		})
	}
}

func given(calc stringcalc.StringCalc) givenFn {
	return func() stringcalc.StringCalc {
		return calc
	}
}

func when(input string) whenFn {
	return func(calc stringcalc.StringCalc) (int, error) {
		return calc.Add(input)
	}
}

func then(want int) thenFn {
	return func(t testing.TB) func(got int, err error) {
		return func(got int, err error) {
			t.Helper()
			ok(t, err)
			equals(t, want, got)
		}
	}
}

func thenFailWith(want error) thenFn {
	return func(t testing.TB) func(got int, err error) {
		return func(got int, err error) {
			t.Helper()
			equals(t, want, err)
		}
	}
}
