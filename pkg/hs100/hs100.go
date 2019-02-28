package hs100

const turnOnCommand = `{"system":{"set_relay_state":{"state":1}}}`
const turnOffCommand = `{"system":{"set_relay_state":{"state":0}}}`

type Hs100 struct {
	Address string
	commandSender CommandSender
}

func NewHs100(address string, s CommandSender) *Hs100 {
	return &Hs100{
		Address: address,
		commandSender: s,
	}
}

type CommandSender interface {
	SendCommand(address string, command string) error
}

func (hs100 *Hs100) TurnOn() {
	_ = hs100.commandSender.SendCommand(hs100.Address, turnOnCommand)
}

func (hs100 *Hs100) TurnOff() {
	_ = hs100.commandSender.SendCommand(hs100.Address, turnOffCommand)
}