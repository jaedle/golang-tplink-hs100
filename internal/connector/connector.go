package connector

import (
	"bufio"
	"net"

	"github.com/jaedle/golang-tplink-hs100/internal/crypto"
)

func (h Hs100) SendCommand(c Command) error {
	conn, err := net.Dial("tcp", h.IPAddress+":9999")
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(conn)
	writer.Write(crypto.EncryptWithHeader(c.C))
	writer.Flush()
	return nil
}

type Command struct {
	C string
}
type Hs100 struct {
	IPAddress string
}
