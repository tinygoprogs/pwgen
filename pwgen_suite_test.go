package pwgen_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPwgen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pwgen Suite")
}
