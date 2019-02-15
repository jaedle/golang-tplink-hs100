package hs100connector

import (
	"bufio"
	"net"
)

func SendCommand(h Hs100) error {
	conn, err := net.Dial("tcp", h.IPAddress+":9999")
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(conn)
	_, _ = writer.WriteString("expected-command")
	_ = writer.Flush()
	return nil
}

func Encrypt(s string) []byte {
	return []byte(s)
}

func Decrypt(d []byte) string {
	return string(d)
}

type Hs100 struct {
	IPAddress string
}
