package hs100connector_test

import (
	"net"
	"time"

	"github.com/jaedle/golang-tplink-hs100/internal/hs100connector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Connector", func() {
	const t = 1 * time.Second

	It("sends expected command", func() {
		l := startServer()
		defer l.Close()

		requestContent := make(chan string)
		go handleRequest(l, requestContent)

		err := hs100connector.SendCommand(localHostDevice())

		var request string
		select {
		case request = <-requestContent:
			break
		case <-time.After(t):
			Fail("received no return value")
		}

		Expect(request).To(Equal("expected-command"))
		Expect(err).NotTo(HaveOccurred())
	})

	It("fails if cannot connect", func() {
		err := hs100connector.SendCommand(localHostDevice())
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

func handleRequest(l net.Listener, response chan string) {
	conn, err := l.Accept()
	if err != nil {
		Fail("can not start server")
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		Fail("can not read request")
	}
	received := string(buf[:n])
	print(received)
	_, _ = conn.Write([]byte(""))
	_ = conn.Close()
	response <- received
}

func localHostDevice() hs100connector.Hs100 {
	return hs100connector.Hs100{
		IPAddress: "127.0.0.1",
	}
}
