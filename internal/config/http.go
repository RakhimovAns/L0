package config

import (
	"net"
	"time"
)

type TimeoutConfig struct {
	Connect time.Duration `json:"connect"`
	Read    time.Duration `json:"read"`
	Write   time.Duration `json:"write"`
	Idle    time.Duration `json:"idle"`
}

func HTTPTimeouts() TimeoutConfig {
	return TimeoutConfig{
		Connect: 10 * time.Second,
		Read:    30 * time.Second,
		Write:   30 * time.Second,
		Idle:    90 * time.Second,
	}
}

func AddrHTTP() string {
	return net.JoinHostPort(cfg.HTTPHost, cfg.HTTPPort)
}
