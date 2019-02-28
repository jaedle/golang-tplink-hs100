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

	const readStateCommand = `{"system":{"get_sysinfo":{}}}`
	It("returns on if device is on", func() {
		s := &commandSender{
			response: onResponse(),
		}
		hs100 := hs100.NewHs100(anIpAddress, s)

		on, err := hs100.IsOn()

		Expect(err).NotTo(HaveOccurred())
		assertOneCommandSend(s, anIpAddress, readStateCommand)
		Expect(on).To(Equal(true))
	})

	It("returns off if device is off", func() {
		s := &commandSender{
			response: offResponse(),
		}
		hs100 := hs100.NewHs100(anIpAddress, s)

		on, err := hs100.IsOn()

		Expect(err).NotTo(HaveOccurred())
		assertOneCommandSend(s, anIpAddress, readStateCommand)
		Expect(on).To(Equal(false))
	})

	It("fails on invalid response", func() {
		s := &commandSender{
			response: "{]",
		}
		hs100 := hs100.NewHs100(anIpAddress, s)

		_, err := hs100.IsOn()

		Expect(err).To(HaveOccurred())
	})

})

func onResponse() string {
	return `{  
		   "system":{  
		      "get_sysinfo":{  
		         "err_code":0,
		         "sw_ver":"1.1.4 Build 170417 Rel.145118",
		         "hw_ver":"1.0",
		         "type":"IOT.SMARTPLUGSWITCH",
		         "model":"HS110(EU)",
		         "mac":"AA:BB:CC:DD:EE:FF",
		         "deviceId":"1234567890123456789012345678901234567890",
		         "hwId":"01234567890123456789012345678912",
		         "fwId":"98765432109876543210987654321032",
		         "oemId":"ABCDEFABCDEFABCDEFABCDEFABCDEFAB",
		         "alias":"sample-plug",
		         "dev_name":"Wi-Fi Smart Plug With Energy Monitoring",
		         "icon_hash":"",
		         "relay_state":1,
		         "on_time":0,
		         "active_mode":"schedule",
		         "feature":"TIM:ENE",
		         "updating":0,
		         "rssi":-65,
		         "led_off":0,
		         "latitude":11.123456,
		         "longitude":50.123456
		      }
		   }
		}`
}

func offResponse() string {
	return `{  
		   "system":{  
		      "get_sysinfo":{  
		         "err_code":0,
		         "sw_ver":"1.1.4 Build 170417 Rel.145118",
		         "hw_ver":"1.0",
		         "type":"IOT.SMARTPLUGSWITCH",
		         "model":"HS110(EU)",
		         "mac":"AA:BB:CC:DD:EE:FF",
		         "deviceId":"1234567890123456789012345678901234567890",
		         "hwId":"01234567890123456789012345678912",
		         "fwId":"98765432109876543210987654321032",
		         "oemId":"ABCDEFABCDEFABCDEFABCDEFABCDEFAB",
		         "alias":"sample-plug",
		         "dev_name":"Wi-Fi Smart Plug With Energy Monitoring",
		         "icon_hash":"",
		         "relay_state":0,
		         "on_time":0,
		         "active_mode":"schedule",
		         "feature":"TIM:ENE",
		         "updating":0,
		         "rssi":-65,
		         "led_off":0,
		         "latitude":11.123456,
		         "longitude":50.123456
		      }
		   }
		}`
}

func assertOneCommandSend(s *commandSender, address string, command string) {
	Expect(s.calls).To(Equal(1))
	Expect(s.address).To(Equal(address))
	Expect(s.command).To(Equal(command))
}

type commandSender struct {
	calls    int
	address  string
	command  string
	response string
}

func (c *commandSender) SendCommand(addr string, cmd string) (string, error) {
	c.calls++
	c.address = addr
	c.command = cmd
	return c.response, nil
}
