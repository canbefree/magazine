package main_test

import (
	"os"

	"github.com/canbefree/magazine/vars"
	jww "github.com/spf13/jwalterweatherman"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	It("root path exist", func() {
		jww.DEBUG.Printf("vars.RootPath:%v", vars.RootPath)
		Expect(vars.RootPath + string(os.PathSeparator) + "magazine.go").To(BeAnExistingFile())
	})
})
