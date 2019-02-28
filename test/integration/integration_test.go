package integration_test

import (
	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integration", func() {
	It("fails", func() {
		h := hs100.NewHs100("localhost", configuration.Default())
		isOn, err := h.IsOn()
		Expect(err).NotTo(HaveOccurred())
		Expect(isOn).To(BeFalse())

		h.TurnOn()
		isOn, err = h.IsOn()
		Expect(err).NotTo(HaveOccurred())
		Expect(isOn).To(BeTrue())
	})
})
