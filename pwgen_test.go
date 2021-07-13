package pwgen

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("pwgen", func() {
	Describe("modifyAlphabet", func() {
		Context("tiny alphabets", func() {
			It("should add include", func() {
				Expect(modifyAlphabet("abc", "", "def")).To(Equal("abcdef"))
			})
			It("should remove excludes", func() {
				Expect(modifyAlphabet("abc", "ac", "")).To(Equal("b"))
			})
		})
	})

	Describe("GenPassword", func() {
		Context("non-random numbers", func() {
			var alphabet = "ABC"
			It("should return the first character", func() {
				randomsrc := strings.NewReader("\x00\x00\x00")
				Expect(GenPassword(&GenPasswordConfig{
					RandomReader: randomsrc,
					Len:          3,
					Alphabet:     &alphabet,
				})).To(Equal("AAA"))
			})
			It("should return the second character", func() {
				randomsrc := strings.NewReader("\x01\x01\x01")
				Expect(GenPassword(&GenPasswordConfig{
					RandomReader: randomsrc,
					Len:          3,
					Alphabet:     &alphabet,
				})).To(Equal("BBB"))
			})
		})
	})

	Describe("insertEachAtRandomPos", func() {
		Context("should insert missing characters", func() {
			It("should not be changed, when all are present", func() {
				randomsrc := strings.NewReader("")
				in := []byte("ABC")
				Expect(insertEachAtRandomPos(in, "ABC", randomsrc)).To(Equal(in))
			})
			It("replace a single character", func() {
				randomsrc := strings.NewReader("\x00")
				in := []byte("A")
				Expect(insertEachAtRandomPos(in, "B", randomsrc)).To(Equal([]byte("B")))
			})
			It("should replace all", func() {
				randomsrc := strings.NewReader("\x00\x01\x02")
				Expect(insertEachAtRandomPos([]byte("ABC"), "DEF", randomsrc)).To(Equal([]byte("DEF")))
			})
		})
	})
})
