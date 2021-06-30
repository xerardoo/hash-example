package imt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xerardoo/hash-example/imt"
	"testing"
)

func TestIMT(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IMT Test Suite")
}

var _ = Describe("Hast testing", func() {
	input1 := []byte{12}
	output1 := []byte{24, 108, 90, 204, 81, 189, 102, 126}

	hash1, err := imt.Generate(input1)
	It("result be hashed", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(hash1).Should(Equal(output1))
	})
})
