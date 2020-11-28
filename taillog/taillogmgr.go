package taillog

import (
	"fmt"
	"github.com/Fish-pro/logagent/etcd"
	"time"
)

var tskMgr *TailLogMgr

type TailLogMgr struct {
	logEntry    []*etcd.LogEntry
	taskMap     map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logConf []*etcd.LogEntry) {
	tskMgr = &TailLogMgr{
		logEntry:    logConf,
		taskMap:     make(map[string]*TailTask, 32),
		newConfChan: make(chan []*etcd.LogEntry),
	}
	for index, value := range tskMgr.logEntry {
		fmt.Printf("index:%v,value:%v\n", index, value)
		task := NewTailTask(value.Path, value.Topic)
		mk := fmt.Sprintf("%s_%s", value.Path, value.Topic)
		tskMgr.taskMap[mk] = task
	}
	go tskMgr.run()
}

// 监听自己的chan
func (t *TailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			fmt.Println(newConf)
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.taskMap[mk]
				if !ok {
					// 原来就有
					task := NewTailTask(conf.Path, conf.Topic)
					t.taskMap[mk] = task
				}
			}
			for _, c1 := range t.logEntry {
				isDelete := true
				for _, c2 := range newConf {
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						break
					}
				}
				if isDelete {
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.taskMap[mk].cancelFunc()
				}
			}

		default:
			time.Sleep(time.Second)
		}
	}
}

func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
