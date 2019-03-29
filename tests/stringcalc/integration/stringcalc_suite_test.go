//+build all integration

package integration_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestStringCalc(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "StringCalcSpec")
}

// var definitions to make tests more eloquent.
var (
	Statement = ginkgo.Describe
	Given     = ginkgo.Context
	Then      = ginkgo.It

	//Spec table.D
	Spec  = table.Entry
	Specs = table.DescribeTable
)
