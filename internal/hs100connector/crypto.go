package hs100connector

func Encrypt(s string) []byte {
	key := byte(0xAB)
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b[i] = s[i] ^ key
		key = b[i]
	}

	return b
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
