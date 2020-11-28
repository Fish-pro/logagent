package main

import (
	"context"
	"fmt"
	"github.com/Fish-pro/logagent/config"
	"github.com/Fish-pro/logagent/etcd"
	"github.com/Fish-pro/logagent/getip"
	"gopkg.in/ini.v1"
	"os"
	"time"
)

func main() {
	var conf config.AppConfig
	err := ini.MapTo(&conf, "./config/config.ini")
	if err != nil {
		fmt.Printf("load config file failed error:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("config file:%v\n", conf)

	err = etcd.Init(conf.Etcd.Address, time.Duration(conf.Etcd.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init etcd success")

	ip, err := getip.GetOutBoundIP()
	if err != nil {
		fmt.Printf("get local getip error:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("local getip is:%s\n", ip)

	cKey := fmt.Sprintf(conf.Etcd.Key, ip)
	//test data
	value := `[{"path":"./test/redis.log","topic":"redis_log"},{"path":"./test/web.log","topic":"web_log"},
{"path":"./test/etcd.log","topic":"etcd_log"}]`
	etcd.Client.Put(context.Background(), cKey, value)
}
