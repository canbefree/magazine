package main_test

import (
	"testing"

	_ "github.com/canbefree/magazine/startup"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	jww "github.com/spf13/jwalterweatherman"
)

func TestCc(t *testing.T) {
	BeforeSuite(func() {
		jww.SetStdoutThreshold(jww.LevelTrace)
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cc Suite")
}
