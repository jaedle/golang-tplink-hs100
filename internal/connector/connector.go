package connector

import (
	"bufio"
	"net"

	"github.com/jaedle/golang-tplink-hs100/internal/crypto"
)

func SendCommand(h Hs100) error {
	conn, err := net.Dial("tcp", h.IPAddress+":9999")
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(conn)
	writer.Write(crypto.EncryptWithHeader("{expected: command}}"))
	writer.Flush()
	return nil
}

type Hs100 struct {
	IPAddress string
}
