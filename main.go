package main

import (
	"fmt"
	"github.com/Fish-pro/logagent/config"
	"github.com/Fish-pro/logagent/etcd"
	"github.com/Fish-pro/logagent/kafka"
	"github.com/Fish-pro/logagent/taillog"
	"gopkg.in/ini.v1"
	"os"
	"sync"
	"time"
)

func main() {
	// 0.加载配置文件
	var conf config.AppConfig
	err := ini.MapTo(&conf, "./config/config.ini")
	if err != nil {
		fmt.Printf("load config file failed error:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("config file:%v\n", conf)

	// 1.初始化kafka的连接
	err = kafka.Init([]string{conf.Kafka.Address}, conf.Kafka.MaxSize)
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

	taillog.Init(logEntrys)

	// 排一个哨兵
	wg := sync.WaitGroup{}
	wg.Add(1)
	go etcd.WatchConfig(conf.Etcd.Key, taillog.NewConfChan())
	wg.Wait()
}
