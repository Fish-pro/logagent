package config

import (
	"gopkg.in/ini.v1"
)

type AppConfig struct {
	Kafka KafkaConfig `ini:"kafka"`
	Tail  TailConfig  `ini:"taillog"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TailConfig struct {
	FileName string `ini:"filename"`
}

func New() (*AppConfig, error) {
	conf, err := ini.Load("./config/config.ini")
	if err != nil {
		return nil, err
	}
	return &AppConfig{
		Kafka: KafkaConfig{
			Address: conf.Section("kafka").Key("address").String(),
			Topic:   conf.Section("kafka").Key("topic").String(),
		},
		Tail: TailConfig{FileName: conf.Section("taillog").Key("filename").String()},
	}, nil
}
