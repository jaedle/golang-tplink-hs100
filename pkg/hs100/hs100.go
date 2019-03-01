package hs100

import (
	"encoding/json"
	"github.com/pkg/errors"
)

const turnOnCommand = `{"system":{"set_relay_state":{"state":1}}}`
const turnOffCommand = `{"system":{"set_relay_state":{"state":0}}}`
const isOnCommand = `{"system":{"get_sysinfo":{}}}`
const currentPowerConsumptionCommand = `{"emeter":{"get_realtime":{},"get_vgain_igain":{}}}`

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
	resp, err := hs100.commandSender.SendCommand(hs100.Address, isOnCommand)
	if err != nil {
		return false, err
	}

	err, on := isOn(resp)
	if err != nil {
		return false, err
	}

	return on, nil
}

func (hs100 *Hs100) GetName() (string, error) {
	resp, err := hs100.commandSender.SendCommand(hs100.Address, isOnCommand)

	if err != nil {
		return "", err
	}

	err, name := name(resp)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (hs100 *Hs100) GetCurrentPowerConsumption() (PowerConsumption, error) {
	resp, err := hs100.commandSender.SendCommand(hs100.Address, currentPowerConsumptionCommand)
	if err != nil {
		return PowerConsumption{}, errors.Wrap(err, "Could not read from hs100 device")
	}

	var r powerConsumptionResponse
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return PowerConsumption{}, errors.Wrap(err, "Cannot parse response")
	}

	return PowerConsumption{
		Current: r.Emeter.RealTime.Current,
	}, nil
}

func name(resp string) (error, string) {
	var r response
	err := json.Unmarshal([]byte(resp), &r)
	name := r.System.SystemInfo.Alias
	return err, name
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
			RelayState int    `json:"relay_state"`
			Alias      string `json:"alias"`
		} `json:"get_sysinfo"`
	} `json:"system"`
}

type powerConsumptionResponse struct {
	Emeter struct {
		RealTime struct {
			Current float32 `json:"current"`
		} `json:"get_realtime"`
	} `json:"emeter"`
}

type PowerConsumption struct {
	Current float32
}
