package thread

import (
	"bytes"
	"fmt"
	"github.com/toophy/pangu/help"
	"os"
)

const (
	LogBuffSize = 20 * 1024 * 1024
)

// 场景线程
type LogThread struct {
	Thread

	buffs   bytes.Buffer // 日志总缓冲
	logFile *os.File
}

// 新建场景线程
func New_log_thread(heart_time int64, lay1_time uint64) (*LogThread, error) {
	a := new(LogThread)
	err := a.Init_log_thread(heart_time, lay1_time)
	if err == nil {
		return a, nil
	}
	return nil, err
}

// 初始化场景线程
func (this *LogThread) Init_log_thread(heart_time int64, lay1_time uint64) error {
	err := this.Init_thread(this, Tid_log, "log_thread", heart_time, lay1_time)
	if err == nil {
		this.buffs.Grow(LogBuffSize)

		name := fmt.Sprintf("../file_%d.log", this.Get_thread_id())
		if !help.IsExist(name) {
			os.Create(name)
		}
		file, err := os.OpenFile(name, os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		this.logFile = file
		this.logFile.Seek(0, 2)

		return nil
	}
	return err
}

// 响应线程首次运行
func (this *LogThread) on_first_run() {
	// 处理文件
	evt := &Event_flush_log{}
	evt.Init("", 300)
	this.PostEvent(evt)
}

// 响应线程退出
func (this *LogThread) on_end() {
	this.logFile.Close()
}

// 响应线程运行
func (this *LogThread) on_run() {
}

func (this *LogThread) Add_log(d string) {
	this.buffs.WriteString(d)
}

func (this *LogThread) Flush_log() {

	this.logFile.Write(this.buffs.Bytes())
	this.buffs.Reset()
}
