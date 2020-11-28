package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var Client *clientv3.Client

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化etcd
func Init(address string, timeout time.Duration) error {
	var err error
	Client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{address},
		DialTimeout: timeout,
	})
	if err != nil {
		return err
	}
	return nil
}

// 根据key获取日志配置项
func Getconfig(key string) ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := Client.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var logEntrys []*LogEntry
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntrys)
		if err != nil {
			return nil, err
		}
	}
	return logEntrys, nil
}

func WatchConfig(key string, newConfChan chan<- []*LogEntry) {
	ch := Client.Watch(context.Background(), key)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed,err:%v]n", err)
					continue
				}
				fmt.Printf("get new config:%v\n", newConf)
			}
			newConfChan <- newConf
		}
	}
}
