package main

import (
	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"log"
	"os"
)

func main() {
	h := hs100.NewHs100("192.168.2.100", configuration.Default())

	info, err := h.GetInfo()
	if err != nil {
		log.Print("Error on accessing device")
		os.Exit(1)
	}

	log.Printf("Name of device: %s", info.Name)
}
