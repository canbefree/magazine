package toml_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"cc/utils/config/toml"
	"cc/vars"
)

var _ = Describe("Toml", func() {
	var configer *toml.TomlLoader
	var setting *vars.Setting
	BeforeEach(func() {
		configer = toml.NewTomlLoader()
		setting = &vars.Setting{
			DefaultMagazineSize: 100,
		}
	})
	FIt("config file saved", func() {
		err := configer.SaveConfig("redis", setting)
		Expect(err).To(BeNil())
	})
})
