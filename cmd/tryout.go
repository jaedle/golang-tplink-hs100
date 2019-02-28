package main

import (
	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"time"
)

func main() {

	h := hs100.NewHs100("192.168.2.100", configuration.Default())

	println("Name of device:")
	name, _ := h.GetName()
	println(name)

	time.Sleep(2000 * time.Millisecond)

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
