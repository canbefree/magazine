package main_test

import (
	"cc/vars"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	It("root path exist", func() {
		Expect(vars.RootPath + string(os.PathSeparator) + "main.go").To(BeAnExistingFile())
	})
})
