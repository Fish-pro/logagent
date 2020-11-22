package main

import (
	"fmt"
	"github.com/Fish-pro/logagent/config"
	"github.com/Fish-pro/logagent/kafka"
	"github.com/Fish-pro/logagent/taillog"
	"gopkg.in/ini.v1"
	"os"
	"time"
)

func run(conf *config.AppConfig) {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2.发送到kafka
			kafka.SendToKafka(conf.Kafka.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0.加载配置文件
	//conf, err := config.New()
	//if err != nil {
	//	fmt.Printf("load config file failed error:%v\n", err)
	//	os.Exit(1)
	//}
	var conf config.AppConfig
	err := ini.MapTo(&conf, "./config/config.ini")
	if err != nil {
		fmt.Printf("load config file failed error:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("config file:%v\n", conf)

	// 1.初始化kafka的连接
	err = kafka.Init([]string{conf.Kafka.Address})
	if err != nil {
		fmt.Printf("init kafka failed,error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init kafka success")

	// 2.打开日志文件收集日志
	err = taillog.Init(conf.Tail.FileName)
	if err != nil {
		fmt.Printf("init taillog failed, error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init tail success")

	run(&conf)
}
