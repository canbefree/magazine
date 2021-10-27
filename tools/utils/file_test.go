package utils_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"cc/tools/utils"
)

var _ = Describe("File", func() {
	It("if param is path", func() {
		err := utils.WriteFile("/tmp/not_exist_path/", []byte("str"), os.ModeAppend)
		Expect(err).To(Equal(utils.ErrFileName))
	})

	It("if sub path not exists", func() {
		err := utils.WriteFile("/tmp/not_exist_path/file", []byte("str"), os.ModeAppend)
		Expect(err).To(BeNil())
		os.RemoveAll("/tmp/not_exist_path")
	})

	It("get curr file pwd", func() {
		currfilePath := utils.GetCurrfilePath()
		Expect(filepath.Base(currfilePath)).To(Equal("utils"))
	})
})
