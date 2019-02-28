package main

import (
	"github.com/jaedle/golang-tplink-hs100/internal/connector"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"time"
)

func main() {

	h := hs100.NewHs100("192.168.2.100", &SendCommandConfiguration{})
	println("Is on:")
	b, _ := h.IsOn()
	println(b)

	time.Sleep(2000 * time.Millisecond)

	println("Turning on")
	h.TurnOn()
	println("done")

	time.Sleep(2000 * time.Millisecond)

	println("Is on:")
	b, _ = h.IsOn()
	println(b)

	time.Sleep(2000 * time.Millisecond)

	println("Turning off")
	h.TurnOff()
	println("done")

	time.Sleep(2000 * time.Millisecond)

	println("Is on:")
	b, _ = h.IsOn()
	println(b)
}

type SendCommandConfiguration struct {
}

func (a *SendCommandConfiguration) SendCommand(addr string, command string) (string, error) {
	return connector.SendCommand(addr, command)
}
