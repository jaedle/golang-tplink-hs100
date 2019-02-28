package main

import (
	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"os"
)

func main() {
	h := hs100.NewHs100("192.168.2.100", configuration.Default())

	name, err := h.GetName()
	if err != nil {
		println("Error on accessing device")
		os.Exit(1)
	}

	println("Name of device: " + name)
}
