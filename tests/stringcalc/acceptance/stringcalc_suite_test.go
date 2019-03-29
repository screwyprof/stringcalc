//+build acceptance

package acceptance_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStringCalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringCalcSpec")
}
