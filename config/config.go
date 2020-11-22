package config

type AppConfig struct {
	Kafka KafkaConfig `ini:"kafka"`
	Etcd  EtcdConfig  `ini:"etcd"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
}

type EtcdConfig struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}

//type TailConfig struct {
//	FileName string `ini:"filename"`
//}
//
//func New() (*AppConfig, error) {
//	conf, err := ini.Load("./config/config.ini")
//	if err != nil {
//		return nil, err
//	}
//	return &AppConfig{
//		Kafka: KafkaConfig{
//			Address: conf.Section("kafka").Key("address").String(),
//			Topic:   conf.Section("kafka").Key("topic").String(),
//		},
//		Tail: TailConfig{FileName: conf.Section("taillog").Key("filename").String()},
//	}, nil
//}
