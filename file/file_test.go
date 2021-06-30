package file_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xerardoo/hash-example/file"
	"testing"
)

func TestFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Test Suite")
}

var _ = Describe("File testing", func() {
	input1 := "https://www.learningcontainer.com/wp-content/uploads/2020/04/sample-text-file.txt"
	var newFile file.File
	var limitSize = 100

	_, err := newFile.Fetch(input1)
	It("file be fetched", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	err2 := newFile.Write("./my-file.txt")
	It("file be writed", func() {
		Expect(err2).NotTo(HaveOccurred())
	})

	err3 := newFile.SetLimitSize(limitSize)
	It("file size be limited", func() {
		Expect(err3).NotTo(HaveOccurred())
	})
})
