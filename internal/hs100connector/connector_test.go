package hs100connector_test

import (
	"log"
	"net"
	"time"

	"github.com/jaedle/golang-tplink-hs100/internal/hs100connector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Connector", func() {
	It("returns no error", func() {
		requestContent := make(chan string)
		go startServer(requestContent)

		err := hs100connector.SendCommand(aDevice("127.0.0.1"))

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

})

func startServer(response chan string) {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := l.Accept()
	if err != nil {
		Fail("can not start server")
	}

	buf := make([]byte, 1024)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		Fail("can not read request")
	}
	received := string(buf[:n])
	print(received)
	_, _ = conn.Write([]byte("Message received."))
	_ = conn.Close()
	response <- received
	l.Close()
}

func aDevice(ip string) hs100connector.Hs100 {
	return hs100connector.Hs100{
		IPAddress: ip,
	}
}
