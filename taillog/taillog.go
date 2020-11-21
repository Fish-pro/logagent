package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

// 专门收集日志文件

var (
	tailObj *tail.Tail
)

func Init(fileName string) error {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}

	var err error
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, error:", err)
		return err
	}
	return nil
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
