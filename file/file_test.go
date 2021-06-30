package file_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xerardoo/hash-example/file"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "File Test Suite")
}

var _ = Describe("File testing", func() {
	input1 := "https://www.learningcontainer.com/wp-content/uploads/2020/04/sample-text-file.txt"
	var newFile file.File
	var limitSize float64 = 100
	var path = ".././my-file.txt"
	os.Remove(path)

	doc, err := newFile.Fetch(input1)
	It("file be fetched", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(len(doc)).Should(BeNumerically(">", 0))
	})

	err4 := file.SetLimitSize(limitSize)
	It("file size be limited", func() {
		Expect(err4).NotTo(HaveOccurred())
		Expect(file.MaxFileSize).Should(Equal(limitSize))
	})

	doc, err = newFile.Fetch(input1)
	It("file be fetched", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(len(doc)).Should(BeNumerically(">", 0))
	})

	err2 := newFile.Write(path)
	It("file be writed", func() {
		Expect(err2).NotTo(HaveOccurred())
	})
})
