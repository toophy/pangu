package thread

import (
	"github.com/toophy/pangu/help"
)

// 事件 : 线程投递的日志
type Event_thread_log struct {
	help.Evt_base
	Data string
}

// 事件执行
func (this *Event_thread_log) Exec(home interface{}) bool {
	home.(*WorldThread).Add_log(this.Data)
	return true
}

// 事件 : 线程投递的日志
type Event_flush_log struct {
	help.Evt_base
}

// 事件执行
func (this *Event_flush_log) Exec(home interface{}) bool {
	home.(*WorldThread).Flush_log()

	evt := &Event_flush_log{}
	evt.Init("", 300)
	home.(*WorldThread).PostEvent(evt)

	return true
}
