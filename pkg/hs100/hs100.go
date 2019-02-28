package hs100

import "encoding/json"

const turnOnCommand = `{"system":{"set_relay_state":{"state":1}}}`
const turnOffCommand = `{"system":{"set_relay_state":{"state":0}}}`
const isOnCommand = `{"system":{"get_sysinfo":{}}}`

type Hs100 struct {
	Address       string
	commandSender CommandSender
}

func NewHs100(address string, s CommandSender) *Hs100 {
	return &Hs100{
		Address:       address,
		commandSender: s,
	}
}

type CommandSender interface {
	SendCommand(address string, command string) (string, error)
}

func (hs100 *Hs100) TurnOn() {
	_, _ = hs100.commandSender.SendCommand(hs100.Address, turnOnCommand)
}

func (hs100 *Hs100) TurnOff() {
	_, _ = hs100.commandSender.SendCommand(hs100.Address, turnOffCommand)
}

func (hs100 *Hs100) IsOn() (bool, error) {
	resp, _ := hs100.commandSender.SendCommand(hs100.Address, isOnCommand)
	err, on := isOn(resp)
	if err != nil {
		return false, err
	}

	return on, nil
}

func isOn(s string) (error, bool) {
	var r response
	err := json.Unmarshal([]byte(s), &r)
	on := r.System.SystemInfo.RelayState == 1
	return err, on
}

type response struct {
	System struct {
		SystemInfo struct {
			RelayState int `json:"relay_state"`
		} `json:"get_sysinfo"`
	} `json:"system"`
}
