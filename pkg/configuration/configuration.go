package configuration

import "github.com/jaedle/golang-tplink-hs100/internal/connector"

type DefaultSendCommandImplementation struct {
}

func (a *DefaultSendCommandImplementation) SendCommand(addr string, command string) (string, error) {
	return connector.SendCommand(addr, command)
}

func DefaultConfiguration() *DefaultSendCommandImplementation {
	return &DefaultSendCommandImplementation{}
}
