package toml_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestToml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Toml Suite")
}
