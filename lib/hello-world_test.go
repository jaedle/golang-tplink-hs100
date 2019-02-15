package lib_test

import (
	"github.com/jaedle/golang-tplink-hs100/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HelloWorld", func() {
	It("greets the world", func() {
		s := lib.HelloWorld()
		Expect(s).To(Equal("Hello World!"))
	})
})
