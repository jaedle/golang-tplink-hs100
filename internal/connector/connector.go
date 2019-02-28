package connector

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/jaedle/golang-tplink-hs100/internal/crypto"
)

const devicePort = ":9999"

func SendCommand(address string, command string) (string, error) {
	conn, err := net.Dial("tcp", address+devicePort)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, err = writer.Write(crypto.EncryptWithHeader(command))
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
