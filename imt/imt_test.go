package imt_test

import (
	"encoding/hex"
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
	output2 := []byte{24, 108, 90, 204, 81, 189, 102, 126, 72, 108, 90, 204, 81, 189, 102, 126}
	outputhex1 := "186c5acc51bd667e"
	outputhex2 := "186c5acc51bd667e486c5acc51bd667e"

	hash1, err := imt.Generate(input1)
	It("result be hashed", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(hash1).Should(Equal(output1))
		Expect(hex.EncodeToString(hash1[:])).Should(Equal(outputhex1))
	})

	hash2, err4 := imt.Generate(append(input1, input1...))
	It("input be empty slice", func() {
		Expect(hash2).Should(Equal(output2))
		Expect(err4).NotTo(HaveOccurred())
		Expect(hex.EncodeToString(hash2[:])).Should(Equal(outputhex2))
	})
})
