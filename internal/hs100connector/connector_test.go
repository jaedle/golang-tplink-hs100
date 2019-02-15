package hs100connector_test

import (
	"net"
	"time"

	"github.com/jaedle/golang-tplink-hs100/internal/hs100connector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Connector", func() {
	It("sends expected command", func() {
		l, err := net.Listen("tcp", ":9999")
		if err != nil {
			Fail("could not start test server")
		}
		defer l.Close()

		requestContent := make(chan string)
		go handleRequest(l, requestContent)

		err = hs100connector.SendCommand(aDevice("127.0.0.1"))

		var r string = ""
		select {
		case r = <-requestContent:
			break
		case <-time.After(1 * time.Second):
			Fail("received no return value")
		}

		Expect(err).NotTo(HaveOccurred())
		Expect(r).To(Equal("expected-command"))
	})

	It("fails if cannot connect", func() {
		err := hs100connector.SendCommand(aDevice("127.0.0.1"))
		Expect(err).To(HaveOccurred())
	})

})

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

func aDevice(ip string) hs100connector.Hs100 {
	return hs100connector.Hs100{
		IPAddress: ip,
	}
}
