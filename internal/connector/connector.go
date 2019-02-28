package connector

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/jaedle/golang-tplink-hs100/internal/crypto"
)

const devicePort = ":9999"

func (h Device) SendCommand(c Command) (string, error) {
	conn, err := net.Dial("tcp", h.ipAddress+devicePort)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, err = writer.Write(crypto.EncryptWithHeader(c.c))
	if err != nil {
		return "", err
	}
	writer.Flush()
	response, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("Could not receive response", err)
	}

	return crypto.DecryptWithHeader(response), nil
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
