package main

import (
	"log"
	"os"

	"github.com/jaedle/golang-tplink-hs100/internal/connector"
)

func main() {
	dev := connector.NewDevice("192.168.2.100")
	cmd := connector.NewCommand(`{"system":{"set_relay_state":{"state":1}}}`)

	response, err := dev.SendCommand(cmd)
	if err != nil {
		log.Fatal("could not send command")
		os.Exit(1)
	}

	println("Command was sent")
	println("Response:")
	println(response)

	cmd = connector.NewCommand(`{"system":{"get_sysinfo":{}}}`)

	response, err = dev.SendCommand(cmd)
	if err != nil {
		log.Fatal("could not send command")
		os.Exit(1)
	}

	println("Command was sent")
	println("Response:")
	println(response)
}
