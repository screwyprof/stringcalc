//+build acceptance

package acceptance_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

const expectedOutput = `0
5
7
7
6
An error occurred: invalid input given: non-numbers
3
An error occurred: negative numbers not allowed: -2,-4
2
6
6`

var _ = Describe("StringCalc", func() {
	var session *gexec.Session

	BeforeEach(func() {
		path := build("github.com/screwyprof/stringcalc/cmd")
		session = run(path)
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("exits with status code 0", func() {
		Eventually(session).Should(gexec.Exit(0))
	})

	It("prints results to output", func() {
		Eventually(session).Should(gbytes.Say(expectedOutput))
	})
})

func build(pkgPath string) string {
	path, err := gexec.Build(pkgPath)
	Expect(err).NotTo(HaveOccurred())

	return path
}

func run(path string) *gexec.Session {
	cmd := exec.Command(path)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
