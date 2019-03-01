package hs100_test

import (
	"errors"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hs100", func() {
	const anIpAddress = "192.168.2.1"
	const aDeviceName = "some-device-name"

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

	XIt("should consider the return value of send command", func() {
		//response command is
		//{"system":{"set_relay_state":{"err_code":0}}}
	})

	const readStateCommand = `{"system":{"get_sysinfo":{}}}`
	const readCurrentPowerConsumptionCommand = `{"emeter":{"get_realtime":{},"get_vgain_igain":{}}}`

	Describe("isOn", func() {
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

		It("fails on error on sending the command", func() {
			s := &commandSender{
				error: true,
			}
			hs100 := hs100.NewHs100(anIpAddress, s)

			_, err := hs100.IsOn()

			Expect(err).To(HaveOccurred())
		})
	})

	Describe("GetName", func() {
		It("read the device name", func() {
			s := &commandSender{
				response: reponseWithName(aDeviceName),
			}
			hs100 := hs100.NewHs100(anIpAddress, s)

			name, err := hs100.GetName()

			Expect(err).NotTo(HaveOccurred())
			assertOneCommandSend(s, anIpAddress, readStateCommand)
			Expect(name).To(Equal(aDeviceName))
		})

		It("fails on invalid response", func() {
			s := &commandSender{
				response: "{]",
			}
			hs100 := hs100.NewHs100(anIpAddress, s)

			name, err := hs100.GetName()

			Expect(err).To(HaveOccurred())
			Expect(name).To(Equal(""))
		})

		It("should fail if sending of command fails", func() {
			s := &commandSender{
				error: true,
			}

			hs100 := hs100.NewHs100(anIpAddress, s)
			name, err := hs100.GetName()

			Expect(err).To(HaveOccurred())
			Expect(name).To(Equal(""))
		})
	})

	Describe("currentPowerConsumption", func() {
		It("reads the current power consumption", func() {
			s := &commandSender{
				response: currentPowerConsumptionResponse(
					"1.2345678",
					"230.123456",
				),
			}
			hs100 := hs100.NewHs100(anIpAddress, s)

			powerConsumption, err := hs100.GetCurrentPowerConsumption()

			Expect(err).NotTo(HaveOccurred())
			assertOneCommandSend(s, anIpAddress, readCurrentPowerConsumptionCommand)
			Expect(powerConsumption.Current).To(BeNumerically("~", 1.2345678, 0.001))
			Expect(powerConsumption.Voltage).To(BeNumerically("~", 230.123456, 0.001))
		})

		It("fails if communication with device is not succesful", func() {
			s := &commandSender{
				error: true,
			}
			hs100 := hs100.NewHs100(anIpAddress, s)

			powerConsumption, err := hs100.GetCurrentPowerConsumption()

			Expect(err).To(HaveOccurred())
			Expect(powerConsumption).To(BeZero())
		})

		It("fails if response is invalid", func() {
			s := &commandSender{
				response: "{]",
			}
			hs100 := hs100.NewHs100(anIpAddress, s)

			powerConsumption, err := hs100.GetCurrentPowerConsumption()

			Expect(err).To(HaveOccurred())
			Expect(powerConsumption).To(BeZero())
		})
	})
})

func currentPowerConsumptionResponse(current string, voltage string) string {
	return `{ 
   "emeter":{  
      "get_realtime":{  
         "current":` + current + `,
         "voltage":` + voltage + `,
         "power":284.103008,
         "total":52.859000,
         "err_code":0
      },
      "get_vgain_igain":{  
         "vgain":13290,
         "igain":16887,
         "err_code":0
      }
   }
}`
}

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

func reponseWithName(name string) string {
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
		         "alias":"` + name + `",
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
	Expect(s.calls).To(Equal(1), "no command send to device")
	Expect(s.address).To(Equal(address), "wrong ip address for device")
	Expect(s.command).To(Equal(command), "wrong command sent to device")
}

type commandSender struct {
	calls    int
	address  string
	command  string
	response string
	error    bool
}

func (c *commandSender) SendCommand(addr string, cmd string) (string, error) {
	if c.error {
		return "", errors.New("some error")
	}
	c.calls++
	c.address = addr
	c.command = cmd
	return c.response, nil
}
