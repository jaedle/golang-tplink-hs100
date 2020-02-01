package configuration

import (
	"github.com/jaedle/golang-tplink-hs100/internal/connector"
	"time"
)

const DefaultTimeout = 5 * time.Second

type DefaultSendCommandImplementation struct {
	timeout time.Duration
}

func (a *DefaultSendCommandImplementation) SendCommand(addr string, command string) (string, error) {
	return connector.SendCommand(addr, command, a.timeout)
}

func Default() *DefaultSendCommandImplementation {
	return &DefaultSendCommandImplementation{
		timeout: DefaultTimeout,
	}
}

func (a *DefaultSendCommandImplementation) WithTimeout(timeout time.Duration) {
	a.timeout = timeout
}
