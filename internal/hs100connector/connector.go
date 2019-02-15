package hs100connector

import (
	"bufio"
	"net"
)

func SendCommand(h Hs100) error {
	conn, _ := net.Dial("tcp", h.IPAddress+":9999")
	writer := bufio.NewWriter(conn)
	_, _ = writer.WriteString("expected-command")
	_ = writer.Flush()
	return nil
}

type Hs100 struct {
	IPAddress string
}
