package hs100connector_test

import (
	b64 "encoding/base64"

	"github.com/jaedle/golang-tplink-hs100/internal/hs100connector"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hs100crypto", func() {
	Describe("encrypt / decrypt", func() {
		It("encrypts and decrypts", func() {
			plain := "{command: example}"
			e := hs100connector.Encrypt(plain)
			d := hs100connector.Decrypt(e)

			Expect(plain).To(Equal(d))
		})

		It("encrypts fixtures", func() {
			e := hs100connector.Encrypt(`{"system":{"set_relay_state":{"state":1}}}`)
			expected, _ := b64.StdEncoding.DecodeString("0PKB+Iv/mvfV75S2xaDUi/mc8JHot8Sw0aXA4tijgfKG55P21O7fot+i")
			Expect(e).To(Equal(expected))
		})

		DescribeTable("decrypts fixtures", func(encrypted string, decrypted string) {
			e, _ := b64.StdEncoding.DecodeString(encrypted)
			Expect(hs100connector.Decrypt(e)).To(Equal(decrypted))
		},
			Entry("fixture-1", "0PKB+Iv/mvfV75S2xaDUi/mc8JHot8Sw0aXA4tijgfKG55P21O7fot+i", `{"system":{"set_relay_state":{"state":1}}}`),
			Entry("fixture-2", "0PDSodir37rX9c+0lLbRtMCf7JXmj+GH6MrwnuuH68u2lus=", `{ "system":{ "get_sysinfo":null } }`),
		)
	})
})
