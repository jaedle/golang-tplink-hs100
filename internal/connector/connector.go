package connector

import (
	"bufio"
	"encoding/binary"
	"io"
	"net"

	"github.com/jaedle/golang-tplink-hs100/internal/crypto"
)

const devicePort = ":9999"
const headerLength = 4

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

	response, err := readHeader(conn)
	if err != nil {
		return "", err
	}

	payload, err := readPayload(conn, payloadLength(response))
	if err != nil {
		return "", err
	}

	return crypto.Decrypt(payload), nil
}

func readHeader(conn net.Conn) ([]byte, error) {
	headerReader := io.LimitReader(conn, int64(headerLength))
	var response = make([]byte, headerLength)
	_, err := headerReader.Read(response)
	return response, err
}

func readPayload(conn net.Conn, length uint32) ([]byte, error) {
	payloadReader := io.LimitReader(conn, int64(length))
	var payload = make([]byte, length)
	_, err := payloadReader.Read(payload)
	return payload, err
}

func payloadLength(header []byte) uint32 {
	payloadLength := binary.BigEndian.Uint32(header)
	return payloadLength
}
