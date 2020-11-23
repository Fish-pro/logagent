package taillog

import (
	"fmt"
	"github.com/Fish-pro/logagent/kafka"
	"github.com/hpcloud/tail"
)

// 专门收集日志文件

type TailTask struct {
	Path     string
	Topic    string
	Instance *tail.Tail
}

func NewTailTask(path, topic string) *TailTask {
	tailTask := &TailTask{
		Path:  path,
		Topic: topic,
	}
	tailTask.Init()
	return tailTask
}

func (t *TailTask) Init() error {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}

	var err error
	t.Instance, err = tail.TailFile(t.Path, config)
	if err != nil {
		fmt.Println("tail file failed, error:", err)
		return err
	}
	go t.Run()
	return nil
}

func (t *TailTask) Run() {
	for {
		select {
		case line := <-t.Instance.Lines:
			//kafka.SendToKafka(t.Topic, line.Text)
			kafka.SendToChan(t.Topic, line.Text)
		}
	}
}
