package config

import "time"

type KafkaConfig struct {
	Broker  string
	Topic   string
	GroupID string
	Timeout time.Duration
}

func Kafka() KafkaConfig {
	return KafkaConfig{
		Broker:  cfg.KafkaBroker,
		Topic:   cfg.KafkaTopic,
		GroupID: cfg.KafkaGroupID,
		Timeout: 10 * time.Second,
	}
}
