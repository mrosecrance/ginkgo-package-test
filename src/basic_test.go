package src_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSrc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Src Suite")
}

var _ = Describe("empty test", func(){
	It("return true", func(){
		Expect(true).To(BeTrue())
	})
})