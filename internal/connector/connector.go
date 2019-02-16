package connector

import (
	"bufio"
	"net"

	"github.com/jaedle/golang-tplink-hs100/internal/crypto"
)

const devicePort = ":9999"

func (h Device) SendCommand(c Command) error {
	conn, err := net.Dial("tcp", h.ipAddress+devicePort)
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
type Device struct {
	ipAddress string
}

func NewDevice(ipAddress string) Device {
	return Device{
		ipAddress: ipAddress,
	}
}
