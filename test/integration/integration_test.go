package integration_test

import (
	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integration", func() {
	It("fails", func() {
		h := hs100.NewHs100("127.0.0.1", configuration.Default())
		isOn, err := h.IsOn()
		Expect(err).NotTo(HaveOccurred())
		Expect(isOn).To(BeFalse())

		err = h.TurnOn()
		Expect(err).NotTo(HaveOccurred())

		isOn, err = h.IsOn()
		Expect(err).NotTo(HaveOccurred())
		Expect(isOn).To(BeTrue())

		consumption, err := h.GetCurrentPowerConsumption()
		Expect(err).NotTo(HaveOccurred())
		Expect(consumption).NotTo(BeZero())
	})
})
