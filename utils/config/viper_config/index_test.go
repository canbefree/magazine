package viper_config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/prashantv/gostub"

	"cc/utils/config/viper_config"
	"cc/vars"
)

var _ = Describe("Viper", func() {
	// vars.ConfigPath = "./123123/"
	Describe("save config with map[string]interface{} ", func() {
		var ctrl *gostub.Stubs
		// var s1 = struct {
		// 	Name   string
		// 	Map    map[string]string
		// 	Struct struct{ Name string }
		// }{
		// 	Name:   "_name",
		// 	Map:    map[string]string{"_map": "_map"},
		// 	Struct: struct{ Name string }{Name: "123"},
		// }
		var s1 map[string]interface{} = make(map[string]interface{})
		s1["_name"] = "_name"
		s1["map"] = map[string]string{"_map": "_map"}

		BeforeEach(func() {
			ctrl = gostub.Stub(&vars.ConfigPath, "/tmp")
		})
		AfterEach(func() {
			ctrl.Reset()
		})
		It("save config", func() {
			viperConfig := viper_config.NewViperConfig()
			err := viperConfig.SaveConfig("test", s1)
			Expect(err).To(BeNil())
		})
	})

	Describe("save config with struct ", func() {
		var ctrl *gostub.Stubs
		var s1 = struct {
			Name   string
			Map    map[string]string
			Struct struct{ Name string }
		}{
			Name:   "_name",
			Map:    map[string]string{"_map": "_map"},
			Struct: struct{ Name string }{Name: "123"},
		}
		BeforeEach(func() {
			ctrl = gostub.Stub(&vars.ConfigPath, "/tmp")
		})
		AfterEach(func() {
			ctrl.Reset()
		})
		It("save config", func() {
			viperConfig := viper_config.NewViperConfig()
			err := viperConfig.SaveConfig("test", s1)
			Expect(err).To(BeNil())
		})
	})
})
