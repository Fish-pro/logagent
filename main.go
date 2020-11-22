package main

import (
	"fmt"
	"github.com/Fish-pro/logagent/config"
	"github.com/Fish-pro/logagent/etcd"
	"github.com/Fish-pro/logagent/kafka"
	"gopkg.in/ini.v1"
	"os"
	"time"
)

//func run(conf *config.AppConfig) {
//	// 1.读取日志
//	for {
//		select {
//		case line := <-taillog.ReadChan():
//			// 2.发送到kafka
//			kafka.SendToKafka(conf.Kafka.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}

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

	err = etcd.Init(conf.Etcd.Address, time.Duration(conf.Etcd.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init etcd success")

	// 2.1从etcd获取要拉取的日志项
	logEntrys, err := etcd.Getconfig(conf.Etcd.Key)
	if err != nil {
		fmt.Printf("get entrys failed,error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("get config from etcd success:", logEntrys)
	for index, value := range logEntrys {
		fmt.Printf("index:%v,value:%v\n", index, value)
	}
	// 2.2哨兵监视日志获取项是否有变化

	// 2.打开日志文件收集日志
	//err = taillog.Init(conf.Tail.FileName)
	//if err != nil {
	//	fmt.Printf("init taillog failed, error:%v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println("init tail success")
	//
	//run(&conf)
}
