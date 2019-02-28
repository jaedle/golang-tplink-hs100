package hs100_test

import (
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hs100", func() {
	const anIpAddress = "192.168.2.1"

	It("sends turn on command", func() {
		s := &commandSender{}
		hs100 := hs100.NewHs100(anIpAddress, s)

		hs100.TurnOn()

		const turnOnCommand = `{"system":{"set_relay_state":{"state":1}}}`
		assertOneCommandSend(s, anIpAddress, turnOnCommand)
	})


	It("sends turn off command", func() {
		s := &commandSender{}
		hs100 := hs100.NewHs100(anIpAddress, s)

		hs100.TurnOff()

		const turnOffCommand = `{"system":{"set_relay_state":{"state":0}}}`
		assertOneCommandSend(s, anIpAddress, turnOffCommand)
	})

	It("asks for power state", func() {
		s := &commandSender{}
		hs100 := hs100.NewHs100(anIpAddress, s)

		on := hs100.IsOn()

		const isOnCommand = `{"system":{"get_sysinfo":{}}}`
		assertOneCommandSend(s, anIpAddress, isOnCommand)
		Expect(on).To(Equal(true))
	})

})

func assertOneCommandSend(s *commandSender, address string, command string) {
	Expect(s.calls).To(Equal(1))
	Expect(s.address).To(Equal(address))
	Expect(s.command).To(Equal(command))
}

type commandSender struct {
	calls   int
	address string
	command string
}

func (c *commandSender) SendCommand(addr string, cmd string) error {
	c.calls++
	c.address = addr
	c.command = cmd
	return nil
}
