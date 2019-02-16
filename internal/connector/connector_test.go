package connector_test

import (
	"net"
	"time"

	c "github.com/jaedle/golang-tplink-hs100/internal/connector"
	"github.com/jaedle/golang-tplink-hs100/internal/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Connector", func() {
	const t = 1 * time.Second

	It("sends expected command", func() {
		l := startServer()
		defer l.Close()

		requestChan := make(chan []byte)
		go handleRequest(l, requestChan)
		dev := localHostDevice()

		err := dev.SendCommand(command(`{"expected": "command"}}`))

		var request []byte
		select {
		case request = <-requestChan:
			break
		case <-time.After(t):
			Fail("received no return value")
		}

		Expect(request).To(Equal(crypto.EncryptWithHeader(`{"expected": "command"}}`)))
		Expect(err).NotTo(HaveOccurred())
	})

	It("fails if cannot connect", func() {
		dev := localHostDevice()
		err := dev.SendCommand(command(""))
		Expect(err).To(HaveOccurred())
	})
})

func startServer() net.Listener {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		Fail("could not start test server")
	}
	return l
}

func handleRequest(l net.Listener, response chan []byte) {
	conn, err := l.Accept()
	if err != nil {
		Fail("can not start server")
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		Fail("can not read request")
	}
	received := buf[:n]
	_, _ = conn.Write([]byte(""))
	_ = conn.Close()
	response <- received
}

func command(cmd string) c.Command {
	return c.Command{
		C: cmd,
	}
}

func localHostDevice() c.Device {
	return c.NewDevice("127.0.0.1")
}
