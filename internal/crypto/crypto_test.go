package crypto_test

import (
	b64 "encoding/base64"

	"github.com/jaedle/golang-tplink-hs100/internal/crypto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hs100crypto", func() {
	Describe("encrypt / decrypt", func() {
		It("encrypts and decrypts", func() {
			plain := "{command: example}"
			result := crypto.Decrypt(crypto.Encrypt(plain))
			Expect(result).To(Equal(plain))
		})

		DescribeTable("encrypts fixtures", func(plain string, encrypted string) {
			expected, _ := b64.StdEncoding.DecodeString(encrypted)
			Expect(crypto.Encrypt(plain)).To(Equal(expected))
		},
			Entry("fixture-1", `{"system":{"set_relay_state":{"state":1}}}`, "0PKB+Iv/mvfV75S2xaDUi/mc8JHot8Sw0aXA4tijgfKG55P21O7fot+i"),
			Entry("fixture-2", `{ "system":{ "get_sysinfo":null } }`, "0PDSodir37rX9c+0lLbRtMCf7JXmj+GH6MrwnuuH68u2lus="),
		)

		DescribeTable("decrypts fixtures", func(encrypted string, plain string) {
			e, _ := b64.StdEncoding.DecodeString(encrypted)
			Expect(crypto.Decrypt(e)).To(Equal(plain))
		},
			Entry("fixture-1", "0PKB+Iv/mvfV75S2xaDUi/mc8JHot8Sw0aXA4tijgfKG55P21O7fot+i", `{"system":{"set_relay_state":{"state":1}}}`),
			Entry("fixture-2", "0PDSodir37rX9c+0lLbRtMCf7JXmj+GH6MrwnuuH68u2lus=", `{ "system":{ "get_sysinfo":null } }`),
		)
	})

	Describe("encrypt / decrypt with header", func() {
		It("encrypts and decrypts", func() {
			plain := `{ "emeter":{ "get_realtime":null } }`
			result := crypto.DecryptWithHeader(crypto.EncryptWithHeader(plain))
			Expect(result).To(Equal(plain))
		})

		DescribeTable("encrypts fixtures", func(plain string, encrypted string) {
			expected, _ := b64.StdEncoding.DecodeString(encrypted)
			Expect(crypto.EncryptWithHeader(plain)).To(Equal(expected))
		},
			Entry("fixture-1", `{"system":{"set_relay_state":{"state":1}}}`, "AAAAKtDygfiL/5r31e+UtsWg1Iv5nPCR6LfEsNGlwOLYo4HyhueT9tTu36Lfog=="),
			Entry("fixture-2", `{ "system":{ "get_sysinfo":null } }`, "AAAAI9Dw0qHYq9+61/XPtJS20bTAn+yV5o/hh+jK8J7rh+vLtpbr"),
		)

		DescribeTable("decrypts fixtures", func(encrypted string, plain string) {
			e, _ := b64.StdEncoding.DecodeString(encrypted)
			Expect(crypto.DecryptWithHeader(e)).To(Equal(plain))
		},
			Entry("fixture-1", "AAAAKtDygfiL/5r31e+UtsWg1Iv5nPCR6LfEsNGlwOLYo4HyhueT9tTu36Lfog==", `{"system":{"set_relay_state":{"state":1}}}`),
			Entry("fixture-2", "AAAAI9Dw0qHYq9+61/XPtJS20bTAn+yV5o/hh+jK8J7rh+vLtpbr", `{ "system":{ "get_sysinfo":null } }`),
		)
	})
})
