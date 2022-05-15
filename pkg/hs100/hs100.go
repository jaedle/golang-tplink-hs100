package hs100

import (
	"encoding/json"
	"github.com/pkg/errors"
)

const turnOnCommand = `{"system":{"set_relay_state":{"state":1}}}`
const turnOffCommand = `{"system":{"set_relay_state":{"state":0}}}`
const isOnCommand = `{"system":{"get_sysinfo":{}}}`
const currentPowerConsumptionCommand = `{"emeter":{"get_realtime":{},"get_vgain_igain":{}}}`

type Info struct {
	Name     string
	DeviceId string
	on       bool
}

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

func (hs100 *Hs100) TurnOn() error {
	resp, err := hs100.commandSender.SendCommand(hs100.Address, turnOnCommand)
	if err != nil {
		return errors.Wrap(err, "error on sending turn on command for device")
	}

	r, err := parseSetRelayResponse(resp)
	if err != nil {
		return errors.Wrap(err, "Could not parse response from device")
	} else if r.errorOccurred() {
		return errors.New("got non zero exit code from device")
	}

	return nil
}

func parseSetRelayResponse(response string) (setRelayResponse, error) {
	var result setRelayResponse
	err := json.Unmarshal([]byte(response), &result)
	return result, err
}

func (r *setRelayResponse) errorOccurred() bool {
	return r.System.SetRelayState.ErrorCode != 0
}

type setRelayResponse struct {
	System struct {
		SetRelayState struct {
			ErrorCode int `json:"err_code"`
		} `json:"set_relay_state"`
	} `json:"system"`
}

func (hs100 *Hs100) TurnOff() error {
	resp, err := hs100.commandSender.SendCommand(hs100.Address, turnOffCommand)
	if err != nil {
		return errors.Wrap(err, "error on sending turn on command for device")
	}

	r, err := parseSetRelayResponse(resp)
	if err != nil {
		return errors.Wrap(err, "Could not parse response from device")
	} else if r.errorOccurred() {
		return errors.New("got non zero exit code from device")
	}

	return nil
}

func (hs100 *Hs100) IsOn() (bool, error) {
	if info, err := hs100.GetInfo(); err != nil {
		return false, err
	} else {
		return info.on, nil
	}
}

type response struct {
	System struct {
		SystemInfo struct {
			RelayState int    `json:"relay_state"`
			Alias      string `json:"alias"`
			DeviceId   string `json:"deviceId"`
		} `json:"get_sysinfo"`
	} `json:"system"`
}

// Deprecated:
// Use GetInfo() instead

func (hs100 *Hs100) GetName() (string, error) {
	info, err := hs100.GetInfo()
	if err != nil {
		return "", err
	} else {
		return info.Name, nil
	}
}

func toInfo(resp string) (*Info, error) {
	var r response
	err := json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return nil, err
	}

	return &Info{
		Name:     r.System.SystemInfo.Alias,
		DeviceId: r.System.SystemInfo.DeviceId,
		on:       r.System.SystemInfo.RelayState == 1,
	}, nil
}

func (hs100 *Hs100) GetCurrentPowerConsumption() (PowerConsumption, error) {
	resp, err := hs100.commandSender.SendCommand(hs100.Address, currentPowerConsumptionCommand)
	if err != nil {
		return PowerConsumption{}, errors.Wrap(err, "Could not read from hs100 device")
	}
	return powerConsumption(resp)
}

func (hs100 *Hs100) GetInfo() (*Info, error) {
	if resp, err := hs100.commandSender.SendCommand(hs100.Address, isOnCommand); err != nil {
		return nil, err
	} else {
		return toInfo(resp)
	}
}

type PowerConsumption struct {
	Current float32
	Voltage float32
	Power   float32
}

func powerConsumption(resp string) (PowerConsumption, error) {
	var r powerConsumptionResponse
	err := json.Unmarshal([]byte(resp), &r)
	if err != nil {
		return PowerConsumption{}, errors.Wrap(err, "Cannot parse response")
	} else {
		return r.toPowerConsumption(), nil
	}
}

type powerConsumptionResponse struct {
	Emeter struct {
		RealTime struct {
			Current float32 `json:"current"`
			Voltage float32 `json:"voltage"`
			Power   float32 `json:"power"`
		} `json:"get_realtime"`
	} `json:"emeter"`
}

func (r *powerConsumptionResponse) toPowerConsumption() PowerConsumption {
	return PowerConsumption{
		Current: r.Emeter.RealTime.Current,
		Voltage: r.Emeter.RealTime.Voltage,
		Power:   r.Emeter.RealTime.Power,
	}
}
