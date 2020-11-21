package main

import (
	"fmt"
	"github.com/Fish-pro/logagent/kafka"
	"github.com/Fish-pro/logagent/taillog"
	"os"
	"time"
)

func run() {
	// 1.读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2.发送到kafka
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 1.初始化kafka的连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed,error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init kafka success")

	// 2.打开日志文件收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("init taillog failed, error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init tail success")

	run()
}
