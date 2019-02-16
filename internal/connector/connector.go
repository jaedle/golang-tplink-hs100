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
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	writer.Write(crypto.EncryptWithHeader(c.c))
	writer.Flush()
	return nil
}

type Command struct {
	c string
}

func NewCommand(cmd string) Command {
	return Command{
		c: cmd,
	}
}

type Device struct {
	ipAddress string
}

func NewDevice(ipAddress string) Device {
	return Device{
		ipAddress: ipAddress,
	}
}
