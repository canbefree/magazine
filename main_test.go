package main_test

import (
	"os"

	"github.com/canbefree/magazine/vars"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	It("root path exist", func() {
		Expect(vars.RootPath + string(os.PathSeparator) + "main.go").To(BeAnExistingFile())
	})
})
