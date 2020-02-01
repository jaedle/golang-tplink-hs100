package hs100

import (
	"github.com/pkg/errors"
	"net"
	"sync"
)

func Discover(subnet string, s CommandSender) ([]*Hs100, error) {
	ips, err := getIpAddresses(subnet)
	if err != nil {
		return nil, err
	}

	result := &discoverResult{
		devices: make([]*Hs100, 0),
		Mutex:   sync.Mutex{},
	}
	var wg sync.WaitGroup
	wg.Add(len(ips))
	for _, current := range ips {
		go tryIp(s, result, &wg, current)
	}
	wg.Wait()

	return result.devices, nil
}

func tryIp(s CommandSender, r *discoverResult, wg *sync.WaitGroup, ip string) {
	defer wg.Done()

	hs100 := NewHs100(ip, s)
	_, err := hs100.GetName()
	if err != nil {
		return
	}

	r.Lock()
	defer r.Unlock()

	r.devices = append(r.devices, hs100)
}

type discoverResult struct {
	devices []*Hs100
	sync.Mutex
}

func getIpAddresses(subnet string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, errors.Wrap(err, "invalid subnet specfied")
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1], nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j > 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
