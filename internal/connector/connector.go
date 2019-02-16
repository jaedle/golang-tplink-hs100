package connector

import (
	"bufio"
	"net"

	"github.com/jaedle/golang-tplink-hs100/internal/crypto"
)

func (h Device) SendCommand(c Command) error {
	conn, err := net.Dial("tcp", h.ipAddress+":9999")
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
