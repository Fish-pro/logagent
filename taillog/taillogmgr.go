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
		// 3.收集日志，发往kafka
		NewTailTask(value.Path, value.Topic)
	}
	go tskMgr.run()
}

// 监听自己的chan
func (t *TailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			fmt.Println(newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
