package utils_test

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/canbefree/magazine/utils"
)

var _ = Describe("File", func() {
	It("if param is path", func() {
		err := utils.WriteFile("/tmp/not_exist_path/", []byte("str"), os.ModePerm)
		Expect(err).To(Equal(utils.ErrFileName))
	})

	It("if sub path not exists", func() {
		err := utils.WriteFile("/tmp/not_exist_path/file", []byte("str"), os.ModePerm)
		Expect(err).To(BeNil())
	})

	It("get curr file pwd", func() {
		currfilePath := utils.GetCurrfilePath()
		Expect(filepath.Base(currfilePath)).To(Equal("utils"))
	})
	It("get caller code path", func() {
		fileName := utils.GetCallerCodeDir(0)
		fmt.Println(fileName)
	})
	It("list dir [recurrence]", func() {
		files := utils.ListDirRecurrence(".", func(s string) bool { return false })
		Expect(len(files)).To(Equal(0))
	})

	FIt("list dir [recurrence]", func() {
		files := utils.ListDirRecurrence(".", func(s string) bool { return true })
		Expect(len(files)).To(BeNumerically(">", 0))
	})
})
