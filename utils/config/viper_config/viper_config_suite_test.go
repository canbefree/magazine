package viper_config_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestViperConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ViperConfig Suite")
}
