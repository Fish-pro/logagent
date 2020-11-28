package taillog

import (
	"context"
	"fmt"
	"github.com/Fish-pro/logagent/kafka"
	"github.com/hpcloud/tail"
)

// 专门收集日志文件

type TailTask struct {
	Path       string
	Topic      string
	Instance   *tail.Tail
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) *TailTask {
	ctx, cancel := context.WithCancel(context.Background())
	tailTask := &TailTask{
		Path:       path,
		Topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
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
		case <-t.ctx.Done():
			fmt.Printf("tail task:%v 结束了\n", fmt.Sprintf("%s_%s", t.Path, t.Topic))
			return
		case line := <-t.Instance.Lines:
			//kafka.SendToKafka(t.Topic, line.Text)
			kafka.SendToChan(t.Topic, line.Text)
		}
	}
}
