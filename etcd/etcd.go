package etcd

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var client *clientv3.Client

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化etcd
func Init(address string, timeout time.Duration) error {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{address},
		DialTimeout: timeout,
	})
	if err != nil {
		return err
	}
	// test data
	//value := `[{"path":"/tmp/redis.log","topic":"redis_log"},{"path":"/tmp/web.log","topic":"web_log"}]`
	//client.Put(context.Background(), "/logagent/collect_config", value)
	return nil
}

// 根据key获取日志配置项
func Getconfig(key string) ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Get(ctx, key)
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
