//+build integration

package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/screwyprof/stringcalc"
)

var _ = Statement("StringCalc", func() {
	var stringCalc stringcalc.StringCalc

	Given("Given empty input", func() {
		When("method Add is called", func() {
			Then("it should return 0", func() {
				Ω(stringCalc.Add("")).Should(BeZero())
			})
		})
	})

	Given("given a few delimiters", func() {
		When("method Add is called", func() {
			actual, err := stringCalc.Add("//[*][%]\n1*2%3")
			Then("it should return the sum", func() {
				Ω(err).ShouldNot(HaveOccurred())
				Ω(actual).Should(BeEquivalentTo(6))
			})
		})
	})

	var _ = Specs("Positive Statements",
		func(input string, expected int) {
			Ω(stringCalc.Add(input)).Should(BeEquivalentTo(expected))
		},
		Spec("ShouldReturnZeroWhenEmptyInputIsGiven", "", 0),
		Spec("ShouldReturnTheSumWhenACustomDelimiterGiven", "//;\n1;2", 3),
	)
})
