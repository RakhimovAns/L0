package config

import "time"

func RedisTimeouts() TimeoutConfig {
	return TimeoutConfig{
		Connect: 5 * time.Second,
		Read:    10 * time.Second,
		Write:   10 * time.Second,
		Idle:    60 * time.Second,
	}
}

func RedisHost() string {
	return cfg.RedisHost
}

func RedisPort() string {
	return cfg.RedisPort
}

func RedisPassword() string {
	return cfg.RedisPassword
}

func RedisDB() int {
	return cfg.RedisDB
}
