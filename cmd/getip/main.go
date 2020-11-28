package main

import (
	"fmt"
	"github.com/Fish-pro/logagent/getip"
	"os"
)

func main() {
	ip, err := getip.GetOutBoundIP()
	if err != nil {
		fmt.Printf("get local getip error:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("local getip is:%s\n", ip)
}
