package crypto

import (
	"encoding/binary"
)

const lengthHeader = 4

func Encrypt(s string) []byte {
	key := byte(0xAB)
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b[i] = s[i] ^ key
		key = b[i]
	}
	return b
}

func EncryptWithHeader(s string) []byte {
	lengthPayload := len(s)
	b := make([]byte, lengthHeader+lengthPayload)
	copy(b[:lengthHeader], header(lengthPayload))
	copy(b[lengthHeader:], Encrypt(s))
	return b
}

func header(lengthPayload int) []byte {
	h := make([]byte, lengthHeader)
	binary.BigEndian.PutUint32(h, uint32(lengthPayload))
	return h
}

func Decrypt(b []byte) string {
	k := byte(0xAB)
	var newKey byte
	for i := 0; i < len(b); i++ {
		newKey = b[i]
		b[i] = b[i] ^ k
		k = newKey
	}

	return string(b)
}

func DecryptWithHeader(b []byte) string {
	return Decrypt(payload(b))
}

func payload(b []byte) []byte {
	return b[lengthHeader:]
}
