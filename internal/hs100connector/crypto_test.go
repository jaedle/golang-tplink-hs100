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
			result := hs100connector.Decrypt(hs100connector.Encrypt(plain))
			Expect(result).To(Equal(plain))
		})

		DescribeTable("encrypts fixtures", func(plain string, encrypted string) {
			expected, _ := b64.StdEncoding.DecodeString(encrypted)
			Expect(hs100connector.Encrypt(plain)).To(Equal(expected))
		},
			Entry("fixture-1", `{"system":{"set_relay_state":{"state":1}}}`, "0PKB+Iv/mvfV75S2xaDUi/mc8JHot8Sw0aXA4tijgfKG55P21O7fot+i"),
			Entry("fixture-2", `{ "system":{ "get_sysinfo":null } }`, "0PDSodir37rX9c+0lLbRtMCf7JXmj+GH6MrwnuuH68u2lus="),
		)

		DescribeTable("decrypts fixtures", func(encrypted string, plain string) {
			e, _ := b64.StdEncoding.DecodeString(encrypted)
			Expect(hs100connector.Decrypt(e)).To(Equal(plain))
		},
			Entry("fixture-1", "0PKB+Iv/mvfV75S2xaDUi/mc8JHot8Sw0aXA4tijgfKG55P21O7fot+i", `{"system":{"set_relay_state":{"state":1}}}`),
			Entry("fixture-2", "0PDSodir37rX9c+0lLbRtMCf7JXmj+GH6MrwnuuH68u2lus=", `{ "system":{ "get_sysinfo":null } }`),
		)
	})
})
