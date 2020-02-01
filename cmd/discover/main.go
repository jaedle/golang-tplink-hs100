package main

import (
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"log"
)

func main() {
	devices, err := hs100.Discover("192.168.2.0/24")
	if err != nil {
		panic(err)
	}

	log.Printf("Found devices: %d", len(devices))
	for _, d := range devices {
		name, _ := d.GetName()
		log.Printf("Device name: %s", name)
	}
}
